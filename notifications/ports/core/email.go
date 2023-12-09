package core

import (
	"context"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notification "github.com/codeclout/AccountEd/notifications/notification-types"
)

type EmailCorePort interface {
	ProcessEmailValidation(ctx context.Context, out notification.ValidateEmailOut) (*pb.ValidateEmailAddressResponse, error)
}
