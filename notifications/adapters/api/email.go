package api

import (
	"context"

	"github.com/pkg/errors"

	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	"github.com/codeclout/AccountEd/pkg/monitoring"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notificatonCore "github.com/codeclout/AccountEd/notifications/ports/core"
	"github.com/codeclout/AccountEd/notifications/ports/framework/driven"
)

type ValidateEmailIn = notifications.ValidateEmailIn

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

func (a *Adapter) ValidateEmailAddress(ctx context.Context, in *ValidateEmailIn, ch chan *pb.ValidateEmailAddressResponse, ech chan error) {
	drivenEmail, e := a.drivenEmail.EmailVerificationProcessor(ctx, in)
	if e != nil {
		x := errors.Wrapf(e, "api-ValidateEmailAddress -> drivenEmail.EmailVerificationProcessor(%v)", ctx)
		ech <- x
		return
	}

	core, e := a.core.ProcessEmailValidation(ctx, *drivenEmail)
	if e != nil {
		x := errors.Wrapf(e, "api-ValidateEmailAddress -> core.ProcessEmailValidation(%v)", ctx)
		ech <- x
		return
	}

	ch <- core
}

func (a *Adapter) SendPreRegistrationEmailAPI(ctx context.Context, in *notifications.NoReplyEmailIn, ch chan *pb.NoReplyEmailNotificationResponse, errorch chan error) {
	x, e := a.core.SendPreRegistrationEmailCore(ctx, in)
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
