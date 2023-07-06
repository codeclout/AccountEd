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
	mpb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type Adapter struct {
	config       map[string]interface{}
	core         core.HomeschoolCore
	grpcProtocol *protocols.ClientAdapter
	log          *slog.Logger
}

// NewAdapter is a constructor function for the Adapter struct. It initializes a new Adapter object and returns its pointer.
// It accepts four parameters:
// config is a map of string keys to interface{} values that stores the configuration parameters for the Adapter.
// core is the core Homeschool interface, which is an implementation of the core domain behavior in the adapter.
// grpc is the ClientAdapter of the gRPC protocol, which enables communication with other services.
// log is a Logger object used to log activity, errors and other useful information during the program's runtime.
func NewAdapter(config map[string]interface{}, core core.HomeschoolCore, grpc *protocols.ClientAdapter, log *slog.Logger) *Adapter {
	return &Adapter{
		config:       config,
		core:         core,
		grpcProtocol: grpc,
		log:          log,
	}
}

// encryptSessionId generates a hashed session ID using the provided session ID and a secret string fetched from AWS resources.
// It first fetches a parameter from SSM, then retrieves a secret value from Secrets Manager, and then creates the hashed session ID
// using the SHA-256 hashing algorithm. Returns the hashed ID as a string and an error in case of any failure.
func (a *Adapter) encryptSessionID(ctx context.Context, id string) (string, error) {
	client := *a.grpcProtocol.MemberClient
	payload := mpb.EncryptedStringRequest{
		SessionId: id,
	}

	encryptedSessionID, e := client.GetEncryptedSessionId(ctx, &payload)
	if e != nil {
		a.log.Error(e.Error())
		return "", errors.Wrap(e, "failed to encrypt session ID")
	}

	return encryptedSessionID.GetEncryptedSessionId(), nil
}

func (a *Adapter) logError(ctx context.Context, msg string) {

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
			"transaction_id", fmt.Sprintf("%x", ctx.Value(memberTypes.LogLabel("transaction_id"))))
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
			"transaction_id", fmt.Sprintf("%x", ctx.Value(memberTypes.LogLabel("transaction_id"))))
		ech <- memberTypes.ErrorCoreDataInvalid(errors.New("0 data for transaction_id:" + *data.Username))
		return
	}

	out, e := a.core.PreRegister(ctx, coreData)
	if e != nil {
		a.log.Error("core -> pre registration: "+*data.Username,
			"request_id", ctx.Value(memberTypes.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(memberTypes.LogLabel("transaction_id"))))
		ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
		return
	}

	hashedSessionID, e := a.encryptSessionID(ctx, out.SessionID)
	out.SessionID = hashedSessionID

	if out.UsernamePending {
		// send auto correct & confirm email address on front end
	}

	if out.RegistrationPending {
		// send a verification email containing the session id in the url
		// capture ip for session
	}

	// otherwise -> asynchronously
	// create and register account

	ch <- out
}
