package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/credentials"

	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type cc = context.Context
type dBStorageClient = dynamov1.DynamoDBStorageServiceClient
type storeMetadata = sessionTypes.SessionStoreMetadata
type sessionEncOut = sessionTypes.TokenCreateOut

type dbClient = dynamov1.DynamoDBStorageServiceClient

type CredentialsAWSPort interface {
	AssumeRoleCredentials(ctx cc, arn, region *string) (*credentials.StaticCredentialsProvider, error)
	GetSessionIdKey(ctx context.Context, awsconfig []byte) (*string, error)
	GetToken(ctx cc, awsconfig []byte, in string, client *dbClient) (*dynamov1.FetchTokenResponse, error)
	StoreToken(ctx cc, client *dBStorageClient, in *sessionEncOut, hasAutoCorrect bool, staticCredentials []byte) error

	/* GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client
	GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client
	GetDynamoClient(ctx context.Context, config *aws.Config, region *string) (*dynamodb.Client, error)
	GetR2StorageClient(ctx context.Context, config *aws.Config, cloudflareAccountID *string) (*s3.Client, error)
	GetSESClient(ctx context.Context, config *aws.Config) *sesv2.Client */
}
