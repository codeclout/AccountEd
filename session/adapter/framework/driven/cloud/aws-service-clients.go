package cloud

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

type ErrorDefaultConfiguration error
type ErrorCredentialsRetrieval error
type ErrorInvalidConfiguration error

type Adapter struct {
	config    map[string]interface{}
	log       *slog.Logger
	timeStamp func() time.Time
}

func NewAdapter(config map[string]interface{}, ts func() time.Time, log *slog.Logger) *Adapter {
	return &Adapter{
		config:    config,
		log:       log,
		timeStamp: ts,
	}
}

func (a *Adapter) AssumeRoleCredentials(ctx context.Context, arn, region *string) (*credentials.StaticCredentialsProvider, error) {
	defaultConfiguration, e := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(*region))
	if e != nil {
		a.log.Error(ErrorDefaultConfiguration(errors.New(e.Error())).Error())
		return nil, ErrorDefaultConfiguration(errors.New("unable to load AWS configuration"))
	}

	client := sts.NewFromConfig(defaultConfiguration)

	stsRoleOutput, e := client.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         arn,
		RoleSessionName: aws.String("session-service" + strconv.Itoa(a.timeStamp().Nanosecond())),
	})
	if e != nil {
		a.log.Error(ErrorDefaultConfiguration(errors.New(e.Error())).Error())
		return nil, ErrorDefaultConfiguration(fmt.Errorf("failed to assume role: %w", e))
	}

	out := credentials.StaticCredentialsProvider{Value: aws.Credentials{
		AccessKeyID:     *stsRoleOutput.Credentials.AccessKeyId,
		SecretAccessKey: *stsRoleOutput.Credentials.SecretAccessKey,
		SessionToken:    *stsRoleOutput.Credentials.SessionToken,
	}}

	return &out, nil
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

func (a *Adapter) GetSESClient(ctx context.Context, config *aws.Config) *sesv2.Client {
	client := sesv2.NewFromConfig(*config)

	// @TODO - cache client
	// @TODO - store and check client expiration
	return client
}
