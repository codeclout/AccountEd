package core

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) ProcessEmailValidation(ctx context.Context) (*notifications.EmailDrivenIn, error) {
	email := ctx.Value(notifications.EmailAddress("address"))
	emailAddress, ok := email.(string)
	if !ok {
		a.monitor.LogGenericError("core -> email address is not a string")
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailAddress"))
	}

	emailProcessorDomain, ok := a.config["EmailProcessorDomain"].(string)
	if !ok {
		a.monitor.LogGenericError("core -> email processor domain is not a string")
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New("core -> wrong type: emailProcessorDomain"))
	}

	emailProcessorPath, ok := a.config["EmailVerifierApiPath"].(string)
	if !ok {
		a.monitor.LogGenericError("core -> email processor path is not a string")
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
