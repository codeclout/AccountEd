package core

import (
	"context"

	emailv1 "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notification "github.com/codeclout/AccountEd/notifications/notification-types"
)

type EmailCorePort interface {
	SendPreRegistrationEmailCore(ctx context.Context) (*notification.NoReplyEmailInput, error)
	ProcessEmailValidation(ctx context.Context, out notification.ValidateEmailOut) (*emailv1.ValidateEmailAddressResponse, error)
}
