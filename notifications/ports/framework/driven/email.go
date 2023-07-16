package driven

import (
  "context"

  notifications "github.com/codeclout/AccountEd/notifications/notification-types"
)

type EmailDrivenPort interface {
  EmailVerificationProcessor(ctx context.Context, in *notifications.EmailDrivenIn) (*notifications.ValidateEmailOut, error)
  SendPreRegistrationEmail(ctx context.Context, awsconfig []byte, body, subject string, in *notifications.NoReplyEmailIn) (*notifications.NoReplyEmailOut, error)
}
