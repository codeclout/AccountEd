package api

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/core"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/framework/driven"
)

type Adapter struct {
	core        core.EmailCorePort
	drivenEmail driven.EmailDrivenPort
	log         *slog.Logger
}

func NewAdapter(log *slog.Logger, core core.EmailCorePort, email driven.EmailDrivenPort) *Adapter {
	return &Adapter{
		core:        core,
		drivenEmail: email,
		log:         log,
	}
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, errorch chan error) {
	ctx = context.WithValue(ctx, "address", address)
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
