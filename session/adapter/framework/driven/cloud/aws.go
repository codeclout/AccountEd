package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type ErrorDefaultConfiguration error
type ErrorCredentialsRetrieval error

type cc = context.Context
type coreMemberSession = sessionTypes.TokenCreateOut

type DBClient = dynamov1.DynamoDBStorageServiceClient

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) createSecretsManagerClient(awsconfig []byte, awsRegion string) (*secretsmanager.Client, error) {
	var creds credentials.StaticCredentialsProvider

	err := json.Unmarshal(awsconfig, &creds)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling awsconfig: %w", err)
	}

	secretsClient := secretsmanager.NewFromConfig(aws.Config{Credentials: creds}, func(options *secretsmanager.Options) {
		options.Region = awsRegion
	})

	return secretsClient, nil
}

func (a *Adapter) createSystemsManagerClient(awsconfig []byte, awsRegion string) (*ssm.Client, error) {
	var creds credentials.StaticCredentialsProvider

	err := json.Unmarshal(awsconfig, &creds)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling awsconfig: %w", err)
	}

	ssmClient := ssm.NewFromConfig(aws.Config{Credentials: creds}, func(options *ssm.Options) {
		options.Region = awsRegion
	})

	return ssmClient, nil
}

func (a *Adapter) getRegion() (string, error) {
	awsRegion, ok := a.config["Region"].(string)
	if !ok {
		return "", errors.New("Check the 'Region' in application configuration")
	}
	return awsRegion, nil
}

func (a *Adapter) getSessionTableName() (*string, error) {
	tableName, ok := a.config["SessionTableName"].(string)

	if !ok {
		const msg = "Invalid configuration: 'SessionTableName' is expected to be a string"
		return nil, errors.New(msg)
	}

	return &tableName, nil
}

func (a *Adapter) AssumeRoleCredentials(ctx context.Context, arn, region *string) (*credentials.StaticCredentialsProvider, error) {
	defaultConfiguration, e := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(*region))
	if e != nil {
		a.monitor.LogGrpcError(ctx, ErrorDefaultConfiguration(e).Error())
		return nil, ErrorDefaultConfiguration(errors.New("unable to load AWS configuration"))
	}

	client := sts.NewFromConfig(defaultConfiguration)

	stsRoleOutput, e := client.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         arn,
		RoleSessionName: aws.String("session-service" + strconv.Itoa(a.monitor.GetTimeStamp().Nanosecond())),
	})
	if e != nil {
		a.monitor.LogGrpcError(ctx, ErrorDefaultConfiguration(e).Error())
		return nil, ErrorDefaultConfiguration(fmt.Errorf("failed to assume role: %w", e))
	}

	out := credentials.StaticCredentialsProvider{Value: aws.Credentials{
		AccessKeyID:     *stsRoleOutput.Credentials.AccessKeyId,
		SecretAccessKey: *stsRoleOutput.Credentials.SecretAccessKey,
		SessionToken:    *stsRoleOutput.Credentials.SessionToken,
	}}

	return &out, nil
}

// GetSessionIdKey retrieves a session ID key from AWS Secret Manager via the System Manager's Parameter Store.
// AWS region and 'PreRegistrationParameter' are extracted from the application configuration.
func (a *Adapter) GetSessionIdKey(ctx context.Context, awsconfig []byte) (*string, error) {
	var creds credentials.StaticCredentialsProvider

	e := json.Unmarshal(awsconfig, &creds)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.InvalidArgument, "Unable to parse AWS credentials. Invalid JSON format")
	}

	awsRegion, ok := a.config["Region"].(string)
	if !ok {
		const msg = "Invalid configuration: 'Region' is expected to be a string"

		a.monitor.LogGrpcError(ctx, msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	ssmClient := ssm.NewFromConfig(aws.Config{Credentials: creds}, func(options *ssm.Options) {
		options.Region = awsRegion
	})
	secretsClient := secretsmanager.NewFromConfig(aws.Config{Credentials: creds}, func(options *secretsmanager.Options) {
		options.Region = awsRegion
	})

	name, ok := a.config["PreRegistrationParameter"].(string)
	if !ok {
		const msg = "Type Error: 'PreRegistrationParameter' value in configuration must be of type string "

		a.monitor.LogGrpcError(ctx, msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	in := ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}

	parameter, e := ssmClient.GetParameter(ctx, &in)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Internal, "Failed to retrieve parameter from SSM")
	}

	parameterValue := parameter.Parameter.Value
	if parameterValue == nil {
		const msg = "Invalid operation: Attempted to access value of a non-existent parameter"

		a.monitor.LogGrpcError(ctx, msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	secIn := secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(*parameterValue),
		VersionStage: aws.String("AWSCURRENT"),
	}

	sec, e := secretsClient.GetSecretValue(ctx, &secIn)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Internal, e.Error())
	}

	return sec.SecretString, nil
}

func (a *Adapter) GetToken(ctx cc, creds []byte, in string, db *DBClient) (*dynamov1.FetchTokenResponse, error) {
	tableName, e := a.getSessionTableName()
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}

	client := *db

	resp, e := client.FetchToken(ctx, &dynamov1.FetchTokenRequest{
		Credentials: creds,
		TableName:   *tableName,
		Token:       in,
	})

	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Unknown, "Invalid operation: Attempt to retrieve session material failed")
	}

	return resp, nil
}

func (a *Adapter) StoreToken(ctx cc, db *DBClient, in *coreMemberSession, hasAutoCorrect bool, staticCredentials []byte) error {
	tableName, e := a.getSessionTableName()
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return status.Error(codes.InvalidArgument, e.Error())
	}

	if db == nil {
		const msg = "Invalid operation: Storage client is nil"

		a.monitor.LogGrpcError(ctx, msg)
		return status.Error(codes.InvalidArgument, msg)
	}

	client := *db

	data := dynamov1.TokenStoreRequest{
		HasAutoCorrect:               hasAutoCorrect,
		MemberId:                     in.MemberID,
		PublicKey:                    in.PublicEncoded,
		SessionServiceAWScredentials: staticCredentials,
		SessionTableName:             *tableName,
		Token:                        in.Token,
		TokenId:                      in.ID,
		Ttl:                          int32((time.Now().Add(in.TTL).Unix())),
	}

	// @TODO - Handle system event from response data -> new member pre confirmation
	cr, e := client.StorePublicToken(ctx, &data)

	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return status.Error(codes.Internal, "failed to put item in DynamoDB table."+
			"ensure the input is correct and you have sufficient permissions")
	}

	a.monitor.LogGenericInfo(fmt.Sprintf("%+v", cr))

	return nil
}

// func (a *Adapter) GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client {
// 	client := ssm.NewFromConfig(*config)
//
// 	// @TODO - cache client
// 	// @TODO - store and check client expiration
// 	return client
// }
//
// func (a *Adapter) GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client {
// 	client := secretsmanager.NewFromConfig(*config)
//
// 	// @TODO - cache client
// 	// @TODO - store and check client expiration
// 	return client
// }
//
// func (a *Adapter) GetR2StorageClient(ctx context.Context, config *aws.Config, cloudflareAccountID *string) (*s3.Client, error) {
// 	creds, e := config.Credentials.Retrieve(ctx)
// 	if e != nil {
// 		a.monitor.LogGenericError(e.Error())
// 		return nil, ErrorCredentialsRetrieval(errors.New("unable to retrieve S3 credentials"))
// 	}
//
// 	s3Config, e := awsconfig.LoadDefaultConfig(ctx,
// 		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
// 			func(service, awsregion string, opts ...interface{}) (aws.Endpoint, error) {
// 				return aws.Endpoint{URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", *cloudflareAccountID)}, nil
// 			})),
// 		awsconfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
// 			Value: aws.Credentials{
// 				AccessKeyID:     creds.AccessKeyID,
// 				SecretAccessKey: creds.SecretAccessKey,
// 				SessionToken:    creds.SessionToken,
// 			},
// 		}),
// 	)
// 	if e != nil {
// 		a.monitor.LogGenericError(e.Error())
// 		return nil, ErrorDefaultConfiguration(errors.New("unable to load S3 configuration"))
// 	}
//
// 	client := s3.NewFromConfig(s3Config)
// 	return client, nil
// }
//
// func (a *Adapter) GetSESClient(ctx context.Context, config *aws.Config) *sesv2.Client {
// 	client := sesv2.NewFromConfig(*config)
//
// 	// @TODO - cache client
// 	// @TODO - store and check client expiration
// 	return client
// }
