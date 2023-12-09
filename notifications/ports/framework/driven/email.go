package driven

import (
	"context"

	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
)

type EmailDrivenPort interface {
	EmailVerificationProcessor(ctx context.Context, in *notifications.ValidateEmailIn) (*notifications.ValidateEmailOut, error)
}
