package api

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/core"
	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

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
func (a *Adapter) PreRegisterPrimaryMember(ctx context.Context, data *memberTypes.PrimaryMemberStartRegisterIn, ch chan *memberTypes.PrimaryMemberStartRegisterOut, ech chan error) {
	emailclient := *a.grpcProtocol.Emailclient
	response, e := emailclient.ValidateEmailAddress(ctx, &pb.ValidateEmailAddressRequest{Address: *data.Username})
	if e != nil {
		a.log.Error(*data.Username,
			"request_id", ctx.Value(memberTypes.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(memberTypes.TransactionID("transaction_id"))))
		ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
		return
	}

	coreData := memberTypes.EmailValidationIn{
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

	if coreData == (memberTypes.EmailValidationIn{}) {
		a.log.Error("core -> 0 data returned: "+*data.Username,
			"request_id", ctx.Value(memberTypes.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(memberTypes.TransactionID("transaction_id"))))
		ech <- memberTypes.ErrorCoreDataInvalid(errors.New("0 data for transaction_id:" + *data.Username))
		return
	}

	out, e := a.core.PreRegister(ctx, coreData)
	if e != nil {
		a.log.Error("core -> pre registration: "+*data.Username,
			"request_id", ctx.Value(memberTypes.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(memberTypes.TransactionID("transaction_id"))))
		ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
		return
	}

	if out.RegistrationPending {

	}

	ch <- out
}
