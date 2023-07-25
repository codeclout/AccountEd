package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/credentials"
)

type CredentialsAWSPort interface {
	AssumeRoleCredentials(ctx context.Context, arn, region *string) (*credentials.StaticCredentialsProvider, error)
	/* GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client
	GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client
	GetDynamoClient(ctx context.Context, config *aws.Config, region *string) (*dynamodb.Client, error)
	GetR2StorageClient(ctx context.Context, config *aws.Config, cloudflareAccountID *string) (*s3.Client, error)
	GetSESClient(ctx context.Context, config *aws.Config) *sesv2.Client */
}
