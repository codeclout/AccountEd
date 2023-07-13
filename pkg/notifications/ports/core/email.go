package core

import (
	"context"

	notification "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type EmailCorePort interface {
	SendPreRegistrationEmailCore(ctx context.Context) (*notification.NoReplyEmailInput, error)
	ProcessEmailValidation(ctx context.Context) (*notification.EmailDrivenIn, error)
}
