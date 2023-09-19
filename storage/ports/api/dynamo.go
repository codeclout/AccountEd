package api

import (
	"context"

	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type DynamoDbApiPort interface {
	CreatePublicTokenItem(ctx context.Context, in *storageTypes.TokenStorePayload, ch chan *pb.TokenStoreResponse, ech chan error)
	GetToken(ctx context.Context, in *storageTypes.FetchTokenIn, ch chan *pb.FetchTokenResponse, ech chan error)
}
