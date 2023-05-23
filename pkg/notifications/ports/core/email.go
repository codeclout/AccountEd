package core

import "context"

type NotificationEmailCore interface {
  ProcessEmailValidation(ctx context.Context)
}
