package core

import (
	"context"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notification "github.com/codeclout/AccountEd/notifications/notification-types"
)

type EmailCorePort interface {
	SendPreRegistrationEmailCore(ctx context.Context, in *notification.NoReplyEmailIn) (*notification.NoReplyEmailInput, error)
	ProcessEmailValidation(ctx context.Context, out notification.ValidateEmailOut) (*pb.ValidateEmailAddressResponse, error)
}
