package core

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/google/uuid"

	notification "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type Adapter struct {
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, log *slog.Logger) *Adapter {
	return &Adapter{
		config: config,
		log:    log,
	}
}

// ProcessEmailValidation takes a context and returns an EmailDrivenIn object with the email address, verification endpoint, and a
// generated SessionID, or an error if any issues arise during the process. It retrieves the email address from the context, constructs the endpoint
// using the email_processor_domain and email_verifier_api_path from the Adapter's configuration, and generates a preRegistrationID using the uuid package.
// The returned EmailDrivenIn object contains these values with the EmailAddress, Endpoint, and SessionID fields filled accordingly.
func (a *Adapter) ProcessEmailValidation(ctx context.Context) (*notification.EmailDrivenIn, error) {
	email := ctx.Value(notification.EmailAddress("address"))
	emailAddress, ok := email.(string)
	if !ok {
		a.log.Error("core -> email address is not a string")
		return nil, notification.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailAddress"))
	}

	emailProcessorDomain, ok := a.config["EmailProcessorDomain"].(string)
	if !ok {
		a.log.Error("core -> email processor domain is not a string")
		return nil, notification.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailProcessorDomain"))
	}

	emailProcessorPath, ok := a.config["EmailVerifierApiPath"].(string)
	if !ok {
		a.log.Error("core -> email processor path is not a string")
		return nil, notification.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailProcessorPath"))
	}

	endpoint := fmt.Sprintf("%s%s", emailProcessorDomain, emailProcessorPath)
	preRegistrationID, _ := uuid.NewRandom()

	out := notification.EmailDrivenIn{
		EmailAddress: emailAddress,
		Endpoint:     endpoint,
		SessionID:    &preRegistrationID,
	}

	return &out, nil
}
