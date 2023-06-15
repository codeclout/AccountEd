package cloud

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

type ErrorDefaultConfiguration error
type ErrorCredentialsRetrieval error
type ErrorInvalidConfiguration error

type Adapter struct {
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, log *slog.Logger) *Adapter {
	return &Adapter{
		config: config,
		log:    log,
	}
}

// AssumeRoleCredentials attempts to assume the specified AWS role and returns an AWS Config object with the
// assumed role's credentials, or an error if the operation failed. The function takes a context.Context and an IAM Role
// Amazon Resource Name (ARN) string as arguments, and returns a pointer to an aws.Config object and an error if any.
//
// The ARN parameter is the Amazon Resource Name for the role you want to assume. The context.Context is used for request
// cancellation and timeouts.
func (a *Adapter) AssumeRoleCredentials(ctx context.Context, arn, region *string) (*aws.Config, error) {
	configloader, e := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(*region))
	if e != nil {
		a.log.Error(e.Error())
		return nil, ErrorDefaultConfiguration(errors.New("unable to load AWS configuration"))
	}

	client := sts.NewFromConfig(configloader)
	creds := stscreds.NewAssumeRoleProvider(client, *arn)

	configloader.Credentials = aws.NewCredentialsCache(creds)
	return &configloader, nil
}

// GetSystemsManagerClient creates and returns a new AWS Systems Manager (SSM) client instance using the provided AWS
// configuration. It takes a context.Context and a pointer to an aws.Config as arguments, and returns a
// pointer to an ssm.Client. The context.Context is used for request cancellation and timeouts, while the
// aws.Config should contain the necessary settings and credentials for connecting to the AWS API.
func (a *Adapter) GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client {
	client := ssm.NewFromConfig(*config)

	// @TODO - cache client
	// @TODO - store and check client expiration
	return client
}

// GetSecretsManagerClient creates and returns a new AWS Secrets Manager client instance using the provided AWS
// configuration. The function takes a context.Context and a pointer to an aws.Config as arguments and returns a
// pointer to a secretsmanager.Client. The context.Context is used for request cancellation and timeouts, while the
// aws.Config should contain the necessary settings and credentials for connecting to the AWS API.
func (a *Adapter) GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client {
	client := secretsmanager.NewFromConfig(*config)

	// @TODO - cache client
	// @TODO - store and check client expiration
	return client
}

// GetDynamoClient creates and returns a new DynamoDB client instance using the provided AWS configuration and region. The function
// takes a context.Context, a pointer to an aws.Config, and a pointer to a string representing the AWS region as arguments and returns a
// pointer to a dynamodb.Client and an error if any. The context.Context is used for request cancellation and timeouts, while the aws.Config
// should contain the necessary settings and credentials for connecting to the AWS API. If there is an error retrieving the credentials or
// loading the configuration, appropriate error messages will be logged and returned.
func (a *Adapter) GetDynamoClient(ctx context.Context, config *aws.Config, region *string) (*dynamodb.Client, error) {
	creds, e := config.Credentials.Retrieve(ctx)
	if e != nil {
		a.log.Error(e.Error())
		return nil, ErrorCredentialsRetrieval(errors.New("unable to retrieve DynamoDB credentials"))
	}

	endpoint, ok := a.config["DynamoEndpoint"]
	if !ok {
		a.log.Error("dynamodb endpoint not configured")
		return nil, ErrorInvalidConfiguration(errors.New("configuration error: DynamoEndpoint"))
	}

	dynamoConfig, e := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(*region),
		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, awsregion string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpoint.(string)}, nil
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
		a.log.Error(e.Error())
		return nil, ErrorDefaultConfiguration(errors.New("unable to load DynamoDB configuration"))
	}

	// @TODO - cache client
	// @TODO - store and check client expiration
	client := dynamodb.NewFromConfig(dynamoConfig)
	return client, nil
}

// GetR2StorageClient creates and returns a new Cloudflare R2 Storage client using the provided AWS configuration and Cloudflare Account ID.
// It takes a context.Context, a pointer to an aws.Config, and a pointer to a string representing the Cloudflare Account ID as arguments,
// and returns a pointer to an s3.Client and an error if any. The context.Context is used for request cancellation and timeouts,
// while the aws.Config should contain the necessary settings and credentials for connecting to the AWS API. If there is an error
// retrieving the credentials or loading the configuration, appropriate error messages will be logged and returned.
func (a *Adapter) GetR2StorageClient(ctx context.Context, config *aws.Config, cloudflareAccountID *string) (*s3.Client, error) {
	creds, e := config.Credentials.Retrieve(ctx)
	if e != nil {
		a.log.Error(e.Error())
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
		a.log.Error(e.Error())
		return nil, ErrorDefaultConfiguration(errors.New("unable to load S3 configuration"))
	}

	client := s3.NewFromConfig(s3Config)
	return client, nil
}
