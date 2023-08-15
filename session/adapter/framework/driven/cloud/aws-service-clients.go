package cloud

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type ErrorDefaultConfiguration error
type ErrorCredentialsRetrieval error
type ErrorInvalidConfiguration error

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

func (a *Adapter) GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client {
	client := ssm.NewFromConfig(*config)

	// @TODO - cache client
	// @TODO - store and check client expiration
	return client
}

func (a *Adapter) GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client {
	client := secretsmanager.NewFromConfig(*config)

	// @TODO - cache client
	// @TODO - store and check client expiration
	return client
}

func (a *Adapter) GetR2StorageClient(ctx context.Context, config *aws.Config, cloudflareAccountID *string) (*s3.Client, error) {
	creds, e := config.Credentials.Retrieve(ctx)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, ErrorCredentialsRetrieval(errors.New("unable to retrieve S3 credentials"))
	}

	s3Config, e := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, awsregion string, opts ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", *cloudflareAccountID)}, nil
			})),
		awsconfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     creds.AccessKeyID,
				SecretAccessKey: creds.SecretAccessKey,
				SessionToken:    creds.SessionToken,
			},
		}),
	)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, ErrorDefaultConfiguration(errors.New("unable to load S3 configuration"))
	}

	client := s3.NewFromConfig(s3Config)
	return client, nil
}

func (a *Adapter) GetSESClient(ctx context.Context, config *aws.Config) *sesv2.Client {
	client := sesv2.NewFromConfig(*config)

	// @TODO - cache client
	// @TODO - store and check client expiration
	return client
}
