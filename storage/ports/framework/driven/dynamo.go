package driven

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type TokenStorePayload = storageTypes.TokenStorePayload
type TokenStoreResult = storageTypes.TokenStoreResult

type FetchTokenIn = storageTypes.FetchTokenIn
type FetchTokenResult = storageTypes.FetchTokenResult

type DynamoDbDrivenPort interface {
	GetDynamoClient(ctx context.Context, creds credentials.StaticCredentialsProvider, region *string) (*dynamodb.Client, error)
	GetTokenItem(ctx context.Context, api DynamodbAPI, data FetchTokenIn) (*FetchTokenResult, error)
	StoreToken(ctx context.Context, api DynamodbAPI, data *TokenStorePayload) (*TokenStoreResult, error)
}

type DynamodbAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
