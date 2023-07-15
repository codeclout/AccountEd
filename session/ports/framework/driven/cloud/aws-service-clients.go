package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type CredentialsAWSPort interface {
	AssumeRoleCredentials(ctx context.Context, arn, region *string) (*aws.Config, error)
	GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client
	GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client
	GetDynamoClient(ctx context.Context, config *aws.Config, region *string) (*dynamodb.Client, error)
	GetR2StorageClient(ctx context.Context, config *aws.Config, cloudflareAccountID *string) (*s3.Client, error)
	GetSESClient(ctx context.Context, config *aws.Config) *sesv2.Client
}
