package core

import (
	"context"

	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type TokenStoreResult = storageTypes.TokenStoreResult

type DynamoDbCorePort interface {
	ProcessFetchedToken(ctx context.Context, in storageTypes.FetchTokenResult) (*dynamov1.FetchTokenResponse, error)
	ProcessStoredToken(ctx context.Context, in *TokenStoreResult) (*dynamov1.TokenStoreResponse, error)
}
