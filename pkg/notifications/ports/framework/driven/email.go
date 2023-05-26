package driven

import (
  "bytes"
  "context"
  "mime/multipart"

  notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type EmailDrivenPort interface {
  EmailVerificationProcessor(ctx context.Context, reqBody *bytes.Buffer, writer *multipart.Writer) (*notifications.ValidateEmailOut, error)
}
