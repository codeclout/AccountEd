package drivers

import (
	"context"

	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type DynamoDbDriverPort interface {
	GetPreRegistrationBySessionId(ctx context.Context, request *pb.PreRegistrationConfirmationRequest) (*pb.PreRegistrationConfirmationResponse, error)
}
