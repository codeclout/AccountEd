package core

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	emailv1 "github.com/codeclout/AccountEd/notifications/gen/email/v1"
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

func (a *Adapter) ProcessEmailValidation(ctx context.Context, in notifications.ValidateEmailOut) (*emailv1.ValidateEmailAddressResponse, error) {
	out := emailv1.ValidateEmailAddressResponse{
		Email:             in.Email,
		Autocorrect:       in.Autocorrect,
		Deliverability:    in.Deliverability,
		QualityScore:      in.QualityScore,
		IsValidFormat:     in.IsValidFormat,
		IsFreeEmail:       in.IsFreeEmail,
		IsDisposableEmail: in.IsDisposableEmail,
		IsRoleEmail:       in.IsRoleEmail,
		IsCatchallEmail:   in.IsCatchallEmail,
		IsMxFound:         in.IsMxFound,
		IsSmtpValid:       in.IsSMTPValid,
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
