package api

import (
	"context"

	notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/core"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/framework/driven"
)

type Adapter struct {
	config      map[string]interface{}
	core        core.EmailCorePort
	drivenEmail driven.EmailDrivenPort
	log         *slog.Logger
}

func NewAdapter(config map[string]interface{}, core core.EmailCorePort, email driven.EmailDrivenPort, log *slog.Logger) *Adapter {
	return &Adapter{
		config:      config,
		core:        core,
		drivenEmail: email,
		log:         log,
	}
}

// ValidateEmailAddress takes a context, an email address, a channel for ValidateEmailAddressResponse, and an error channel and validates the email address.
// The results are sent back through the respective channels. If an error occurs during the process, it is sent through the error channel.
func (a *Adapter) ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, errorch chan error) {
	emailAddress := notifications.EmailAddress("address")
	ctx = context.WithValue(ctx, emailAddress, address)

	coreEmailProcessor, e := a.core.ProcessEmailValidation(ctx)
	if e != nil {
		x := errors.Wrapf(e, "api-ValidateEmailAddress -> core.ProcessEmailValidation(%v)", ctx)
		errorch <- x
		return
	}

	validated, e := a.drivenEmail.EmailVerificationProcessor(ctx, coreEmailProcessor)
	if e != nil {
		x := errors.Wrapf(e, "api-ValidateEmailAddress -> drivenEmail.EmailVerificationProcessor(%v)", ctx)
		errorch <- x
		return
	}

	out := pb.ValidateEmailAddressResponse{
		Email:             validated.Email,
		Autocorrect:       validated.Autocorrect,
		Deliverability:    validated.Deliverability,
		QualityScore:      validated.QualityScore,
		IsValidFormat:     validated.IsValidFormat,
		IsFreeEmail:       validated.IsFreeEmail,
		IsDisposableEmail: validated.IsDisposableEmail,
		IsRoleEmail:       validated.IsRoleEmail,
		IsCatchallEmail:   validated.IsCatchallEmail,
		IsMxFound:         validated.IsMxFound,
		IsSmtpValid:       validated.IsSMTPValid,
	}

	ch <- &out
}
