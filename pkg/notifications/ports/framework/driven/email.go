package driven

import (
  "context"

  notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type EmailDrivenPort interface {
  EmailVerificationProcessor(ctx context.Context, in *notifications.EmailDrivenIn) (*notifications.ValidateEmailOut, error)
}
