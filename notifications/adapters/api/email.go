package api

import (
	"context"
	"fmt"

	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"

	"github.com/pkg/errors"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notificatonCore "github.com/codeclout/AccountEd/notifications/ports/core"
	"github.com/codeclout/AccountEd/notifications/ports/framework/driven"
)

type Adapter struct {
	config      map[string]interface{}
	core        notificatonCore.EmailCorePort
	drivenEmail driven.EmailDrivenPort
	monitor     monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, core notificatonCore.EmailCorePort, email driven.EmailDrivenPort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:      config,
		core:        core,
		drivenEmail: email,
		monitor:     monitor,
	}
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, ech chan error) {
	var s string

	if address == (s) {
		const msg = "api => email address is invalid"

		a.monitor.LogGenericError(msg)
		ech <- errors.New(msg)
		return
	}

	emailProcessorDomain, ok := a.config["EmailProcessorDomain"].(string)
	if !ok {
		const msg = "api => email verification domain is invalid"

		a.monitor.LogGenericError(msg)
		ech <- errors.New(msg)
		return
	}

	emailProcessorPath, ok := a.config["EmailVerifierApiPath"].(string)
	if !ok {
		const msg = "api => email verification processor path is invalid"

		a.monitor.LogGenericError(msg)
		ech <- errors.New(msg)
		return
	}

	endpoint := fmt.Sprintf("%s%s", emailProcessorDomain, emailProcessorPath)

	validatorData := notifications.EmailDrivenIn{
		EmailAddress: address,
		Endpoint:     endpoint,
	}

	validated, e := a.drivenEmail.EmailVerificationProcessor(ctx, &validatorData)
	if e != nil {
		x := errors.Wrapf(e, "api-ValidateEmailAddress -> drivenEmail.EmailVerificationProcessor(%v)", ctx)
		ech <- x
		return
	}

	core, e := a.core.ProcessEmailValidation(ctx, *validated)
	if e != nil {
		x := errors.Wrapf(e, "api-ValidateEmailAddress -> core.ProcessEmailValidation(%v)", ctx)
		ech <- x
		return
	}

	ch <- core
}

func (a *Adapter) SendPreRegistrationEmailAPI(ctx context.Context, in *notifications.NoReplyEmailIn, ch chan *pb.NoReplyEmailNotificationResponse, errorch chan error) {
	domain := notifications.URL("domain")
	fromAddress := notifications.EmailAddress("fromAddress")
	sessionID := notifications.SessionID("sessionId")
	toAddress := notifications.EmailAddress("toAddress")

	ctx = context.WithValue(ctx, domain, in.Domain)
	ctx = context.WithValue(ctx, fromAddress, in.FromAddress)
	ctx = context.WithValue(ctx, sessionID, in.SessionID)
	ctx = context.WithValue(ctx, toAddress, in.ToAddress)

	x, e := a.core.SendPreRegistrationEmailCore(ctx)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		errorOut := errors.Wrapf(e, "api-SendPreRegistrationEmailAPI -> core.SendPreRegistrationEmailCore(%v)", ctx)

		errorch <- errorOut
		return
	}

	messageID, e := a.drivenEmail.SendPreRegistrationEmail(ctx, in.AWSCredentials, x.Body, x.Subject, in)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		errorOut := errors.Wrapf(e, "drivenEmail.SendPreRegistrationEmail failed to send email -> (%v)", ctx)

		errorch <- errorOut
		return
	}

	out := pb.NoReplyEmailNotificationResponse{MessageId: messageID.MessageID}

	ch <- &out
}
