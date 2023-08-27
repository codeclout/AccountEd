package driven

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type DynamoDbDrivenPort interface {
	GetDynamoClient(ctx context.Context, creds *credentials.StaticCredentialsProvider, region *string) (*dynamodb.Client, error)
	StoreSession(ctx context.Context, api DynamodbAPI, data storageTypes.PreRegistrationSessionAPIin) (*storageTypes.PreRegistrationSessionDrivenOut, error)
}

type DynamodbAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
