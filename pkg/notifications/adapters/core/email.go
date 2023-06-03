package core

import (
  "context"
  "fmt"

  "github.com/google/uuid"

  notification "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type Adapter struct {
  config map[string]interface{}
}

func NewAdapter(config map[string]interface{}) *Adapter {
  return &Adapter{
    config: config,
  }
}

func (a *Adapter) ProcessEmailValidation(ctx context.Context) (*notification.EmailDrivenIn, error) {
  email := ctx.Value("address").(string)
  endpoint := fmt.Sprintf("%s%s", a.config["email_processor_domain"].(string), a.config["email_verifier_api_path"].(string))

  preRegistrationID, _ := uuid.NewRandom()

  out := notification.EmailDrivenIn{
    EmailAddress: email,
    Endpoint:     endpoint,
    SessionID:    &preRegistrationID,
  }

  return &out, nil
}
