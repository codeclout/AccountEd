package drivers

import (
	"context"

	pb "github.com/codeclout/AccountEd/notifications/gen/v1"
)

type EmailDriverPort interface {
	SendPreRegistrationEmail(context.Context, *pb.NoReplyEmailNotificationRequest) (*pb.NoReplyEmailNotificationResponse, error)
	ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error)
}
