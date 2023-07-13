package api

import (
	"context"
	"fmt"
	"net/url"

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

func NewAdapter(config map[string]interface{}, core core.HomeschoolCore, grpc *protocols.ClientAdapter, log *slog.Logger) *Adapter {
	return &Adapter{
		config:       config,
		core:         core,
		grpcProtocol: grpc,
		log:          log,
	}
}

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

func (a *Adapter) getDomain() (string, error) {
	domain, ok := a.config["Domain"].(string)
	if !ok {
		return domain, errors.New("missing environment variable -> Domain")
	}

	return domain, nil
}

// getEmailDomain is a method of Adapter struct that extracts and returns the domain part of an URL.
func (a *Adapter) getEmailDomain() (string, error) {
	v, e := a.getDomain()
	if e != nil {
		return "", e
	}

	r, _ := url.Parse(v)
	if r == nil || r.Hostname() == "" {
		return "", errors.New("unable to parse hostname")
	}

	return r.Hostname(), nil
}

func (a *Adapter) sendPreRegistrationEmail(ctx context.Context, hashedSessionID string, toAddress []string) error {
	emailclient := *a.grpcProtocol.Emailclient
	fqdn, e := a.getEmailDomain()
	if e != nil {
		return errors.New(e.Error())
	}

	uri, e := a.getDomain()
	if e != nil {
		return errors.New(e.Error())
	}

	reqData := pb.NoReplyEmailNotificationRequest{
		AwsCredentials: a.getAWSCredentialBytes(ctx),
		Domain:         uri,
		FromAddress:    fmt.Sprintf("no-reply@%s", fqdn),
		SessionId:      hashedSessionID,
		ToAddress:      toAddress,
	}

	_, e = emailclient.SendPreRegistrationEmail(ctx, &reqData)
	if e != nil {
		a.log.Error(fmt.Sprintf("pre-registration email failed to process -> %s", e.Error()))
		return errors.Wrapf(e, "domain -> %s pre registration email failed", uri)
	}

	return nil
}

func (a *Adapter) PreRegisterPrimaryMemberAPI(ctx context.Context, data *memberTypes.PrimaryMemberStartRegisterIn, ch chan *memberTypes.PrimaryMemberStartRegisterOut, ech chan error) {
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

	hashedSessionID, _ := a.encryptSessionID(ctx, out.SessionID)
	out.SessionID = hashedSessionID

	if out.UsernamePending {
		// send auto correct & confirm email address on front end
	}

	// if out.RegistrationPending {
	e = a.sendPreRegistrationEmail(ctx, hashedSessionID, []string{response.GetEmail()})

	if e != nil {
		a.log.Error(e.Error())
		ech <- errors.New("unable to process verification email")
	}
	// }

	// otherwise -> asynchronously
	// create and register account

	ch <- out
}
