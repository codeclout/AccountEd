package core

import (
  "bytes"
  "context"
  "mime/multipart"

  "github.com/google/uuid"

  notification "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/framework/driven"
)

type Adapter struct {
  drivenEmail driven.EmailDrivenPort
}

func NewAdapter(drivenEmail driven.EmailDrivenPort) *Adapter {
  return &Adapter{
    drivenEmail: drivenEmail,
  }
}

func (a *Adapter) ProcessEmailValidation(ctx context.Context) (notification.ValidateEmailOut, error) {
  email := ctx.Value("address").(string)
  preRegistrationID, _ := uuid.NewRandom()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  address, _ := writer.CreateFormField("address")

  _, _ = address.Write([]byte(email))
  _ = writer.Close()

  out, e := a.drivenEmail.EmailVerificationProcessor(ctx, body, writer)

  return *out, nil
}
