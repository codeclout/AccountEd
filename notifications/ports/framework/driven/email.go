package driven

import (
	"context"

	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
)

type EmailDrivenPort interface {
	EmailVerificationProcessor(ctx context.Context, in *notifications.ValidateEmailIn) (*notifications.ValidateEmailOut, error)
	SendPreRegistrationEmail(ctx context.Context, awsBytes []byte, body, subject string, in *notifications.NoReplyEmailIn) (*notifications.NoReplyEmailOut, error)
}
