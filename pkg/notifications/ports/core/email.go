package core

import (
  "context"

  notification "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type EmailCorePort interface {
  ProcessEmailValidation(ctx context.Context) (*notification.EmailDrivenIn, error)
}
