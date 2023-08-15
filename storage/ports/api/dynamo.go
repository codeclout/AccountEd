package api

import (
	"context"

	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type preconfirmin = storageTypes.PreRegistrationSessionAPIin
type preconfirmResponsepb = pb.PreRegistrationConfirmationResponse

type DynamoDbApiPort interface {
	PreRegistrationConfirmationApi(ctx context.Context, in preconfirmin, ch chan *preconfirmResponsepb, ech chan error)
}
