package drivers

import (
	"context"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
)

type EmailDriverPort interface {
	ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error)
}
