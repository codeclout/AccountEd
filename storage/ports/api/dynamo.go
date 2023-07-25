package api

import (
	"context"

	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type DynamoDbApiPort interface {
	PreRegistrationConfirmationApi(ctx context.Context, sessionID string, ch chan *pb.PreRegistrationConfirmationResponse, ech chan error)
}
