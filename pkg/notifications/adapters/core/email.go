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

// ProcessEmailValidation takes a context and returns an EmailDrivenIn object with the email address, verification endpoint, and a
// generated SessionID, or an error if any issues arise during the process. It retrieves the email address from the context, constructs the endpoint
// using the email_processor_domain and email_verifier_api_path from the Adapter's configuration, and generates a preRegistrationID using the uuid package.
// The returned EmailDrivenIn object contains these values with the EmailAddress, Endpoint, and SessionID fields filled accordingly.
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
