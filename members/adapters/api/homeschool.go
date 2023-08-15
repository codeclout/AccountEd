package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	sessionpb "github.com/codeclout/AccountEd/session/gen/members/v1"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	membersCore "github.com/codeclout/AccountEd/members/ports/core"
)

type Adapter struct {
	config             map[string]interface{}
	contextAPILabel    memberTypes.ContextAPILabel
	contextDrivenLabel memberTypes.ContextDrivenLabel
	core               membersCore.HomeschoolCore
	gRPC               *protocol.AdapterServiceClients
	monitor            monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, core membersCore.HomeschoolCore, grpc *protocol.AdapterServiceClients, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:             config,
		contextAPILabel:    "api_input",
		contextDrivenLabel: "driven_input",
		core:               core,
		gRPC:               grpc,
		monitor:            monitor,
	}
}

func (a *Adapter) encryptSessionID(ctx context.Context, in *memberTypes.PrimaryMemberStartRegisterOut) (string, error) {
	var s string

	client := *a.gRPC.MemberSessionclient
	payload := sessionpb.EncryptedStringRequest{
		HasAutoCorrect: in.AutoCorrect == (s),
		MemberId:       in.MemberID,
		SessionId:      in.SessionID,
	}

	encryptedSessionID, e := client.GetEncryptedSessionId(ctx, &payload)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
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
	emailclient := *a.gRPC.EmailNotificationclient
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
		a.monitor.LogGenericError(fmt.Sprintf("pre-registration email failed to process -> %s", e.Error()))
		return errors.Wrapf(e, "domain -> %s pre registration email failed", uri)
	}

	return nil
}

func (a *Adapter) PreRegisterPrimaryMember(ctx context.Context, data *memberTypes.PrimaryMemberStartRegisterIn, ch chan *memberTypes.PrimaryMemberStartRegisterOut, ech chan error) {
	emailclient := *a.gRPC.EmailNotificationclient
	response, e := emailclient.ValidateEmailAddress(ctx, &pb.ValidateEmailAddressRequest{Address: *data.Username})
	if e != nil {
		a.monitor.LogHttpError(ctx, *data.Username)
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
		a.monitor.LogHttpError(ctx, "core -> 0 data returned: "+*data.Username)
		ech <- memberTypes.ErrorCoreDataInvalid(errors.New("0 data for transaction_id:" + *data.Username))
		return
	}

	ctx = context.WithValue(ctx, a.contextAPILabel, coreData)
	core, e := a.core.PreRegister(ctx)
	if e != nil {
		a.monitor.LogHttpError(ctx, "core -> pre registration: "+*data.Username)
		ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
		return
	}

	hashedSessionID, _ := a.encryptSessionID(ctx, core)
	core.SessionID = hashedSessionID

	if core.RegistrationPending {
		e = a.sendPreRegistrationEmail(ctx, hashedSessionID, []string{core.MemberID})
		if e != nil {
			a.monitor.LogGenericError(e.Error())
			ech <- errors.New("unable to process verification email")
			return
		}

		ch <- core
	}

	// otherwise -> asynchronously
	// create and register account

	ch <- core
}
