package api

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	mt "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/core"
	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

var ErrorCoreDataInvalid error

type Adapter struct {
	core         core.HomeschoolCore
	grpcProtocol *protocols.ClientAdapter
	log          *slog.Logger
}

func NewAdapter(core core.HomeschoolCore, grpc *protocols.ClientAdapter, log *slog.Logger) *Adapter {
	return &Adapter{
		core:         core,
		grpcProtocol: grpc,
		log:          log,
	}
}

// PreRegisterPrimaryMember is a method of Adapter struct that pre-registers a primary member using provided data.
// It validates the email address, and then passes the validation results and other data to the Core PreRegister
// method. The output is sent to a channel and any errors are sent to an error channel.
func (a *Adapter) PreRegisterPrimaryMember(ctx context.Context, data *mt.PrimaryMemberStartRegisterIn, ch chan *mt.PrimaryMemberStartRegisterOut, ech chan error) {
	emailclient := *a.grpcProtocol.Emailclient
	response, e := emailclient.ValidateEmailAddress(ctx, &pb.ValidateEmailAddressRequest{Address: *data.Username})
	if e != nil {
		ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
		return
	}

	coreData := mt.EmailValidationIn{
		Email:             response.GetEmail(),
		Autocorrect:       response.GetAutocorrect(),
		Deliverability:    response.GetDeliverability(),
		QualityScore:      response.GetQualityScore(),
		IsValidFormat:     response.GetIsValidFormat(),
		IsFreeEmail:       response.GetIsFreeEmail(),
		IsDisposableEmail: response.GetIsDisposableEmail(),
		IsRoleEmail:       response.GetIsRoleEmail(),
		IsCatchallEmail:   response.GetIsCatchallEmail(),
		IsMxFound:         response.GetIsMxFound(),
		IsSmtpValid:       response.GetIsSmtpValid(),
	}

	if coreData == (mt.EmailValidationIn{}) {
		ErrorCoreDataInvalid = errors.New("0 data")
		ech <- ErrorCoreDataInvalid
		return
	}

	out, e := a.core.PreRegister(ctx, coreData)
	if e != nil {
		ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
		return
	}

	ch <- out
}
