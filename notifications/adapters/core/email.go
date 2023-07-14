package core

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/google/uuid"

	notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
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
func (a *Adapter) ProcessEmailValidation(ctx context.Context) (*notifications.EmailDrivenIn, error) {
	email := ctx.Value(notifications.EmailAddress("address"))
	emailAddress, ok := email.(string)
	if !ok {
		a.log.Error("core -> email address is not a string")
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailAddress"))
	}

	emailProcessorDomain, ok := a.config["EmailProcessorDomain"].(string)
	if !ok {
		a.log.Error("core -> email processor domain is not a string")
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailProcessorDomain"))
	}

	emailProcessorPath, ok := a.config["EmailVerifierApiPath"].(string)
	if !ok {
		a.log.Error("core -> email processor path is not a string")
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailProcessorPath"))
	}

	endpoint := fmt.Sprintf("%s%s", emailProcessorDomain, emailProcessorPath)
	preRegistrationID, _ := uuid.NewRandom()

	out := notifications.EmailDrivenIn{
		EmailAddress: emailAddress,
		Endpoint:     endpoint,
		SessionID:    &preRegistrationID,
	}

	return &out, nil
}

func (a *Adapter) SendPreRegistrationEmailCore(ctx context.Context) (*notifications.NoReplyEmailInput, error) {
	domain, ok := ctx.Value(notifications.URL("domain")).(string)
	if !ok {
		return nil, errors.New("incorrect type -> ctx value domain")
	}

	session, ok := ctx.Value(notifications.SessionID("sessionId")).(string)
	if !ok {
		return nil, errors.New("incorrect type -> ctx value sessionId")
	}

	body := "Before you can begin using your new account, it needs to be activated - " +
		"this is a necessary step to verify that the email address you provided is valid and owned by you. " +
		"To proceed with the activation of your account, please click on the link provided below " +
		"or completely copy and paste it into your browser's address bar."
	sessionLink := fmt.Sprintf("%s/member/confirm/%s", domain, session)
	welcome := fmt.Sprintf("Welcom to %s", domain)

	msg := fmt.Sprintf("%s\n\n%s\n\n%s\n\nThanks!\nThe "+domain+" staff", welcome, body, sessionLink)

	out := notifications.NoReplyEmailInput{
		Body:    msg,
		Subject: "Please confirm your email address",
	}

	return &out, nil
}
