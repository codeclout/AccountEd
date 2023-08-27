package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/codeclout/AccountEd/members/ports/framework/driven"
	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	sessionpb "github.com/codeclout/AccountEd/session/gen/members/v1"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	membersCore "github.com/codeclout/AccountEd/members/ports/core"
)

type config = map[string]interface{}
type corePort = membersCore.HomeschoolCore

type drivenPort = driven.HomeschoolDrivenPort
type gRPCclients = protocol.AdapterServiceClients

type pmrStart = memberTypes.PrimaryMemberStartRegisterIn
type pmrOut = memberTypes.PrimaryMemberStartRegisterOut

type validatedEmailIn = memberTypes.VerifiedEmailIn
type validateEmailResponse = pb.ValidateEmailAddressResponse

type Adapter struct {
	config             config
	contextAPILabel    memberTypes.ContextAPILabel
	contextDrivenLabel memberTypes.ContextDrivenLabel
	core               corePort
	driven             drivenPort
	gRPC               *gRPCclients
	monitor            monitoring.Adapter
}

func NewAdapter(config config, core corePort, grpc *gRPCclients, driven drivenPort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:             config,
		contextAPILabel:    "api_input",
		contextDrivenLabel: "driven_input",
		core:               core,
		driven:             driven,
		gRPC:               grpc,
		monitor:            monitor,
	}
}

func (a *Adapter) encryptSessionID(ctx context.Context, in *memberTypes.PrimaryMemberStartRegisterOut) (string, error) {
	if a.gRPC == nil || a.gRPC.MemberSessionclient == nil {
		return "", errors.New("nil gRPC or MemberSessionClient")
	}

	if in == nil {
		return "", errors.New("input parameter is nil")
	}

	payload := sessionpb.EncryptedStringRequest{
		HasAutoCorrect: len(in.AutoCorrect) > 0,
		MemberId:       in.MemberID,
		SessionId:      in.SessionID,
	}

	client := *a.gRPC.MemberSessionclient

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
		const e = "homeschool API missing environment variable -> Domain"

		a.monitor.LogGenericError(e)
		return domain, errors.New(e)
	}

	return domain, nil
}

func (a *Adapter) getEmailDomain() (string, error) {
	var s string

	v, e := a.getDomain()
	if e != nil {
		return "", e
	}

	r, e := url.Parse(v)
	if e != nil || r.Hostname() == (s) {
		a.monitor.LogGenericError("failed to parse email domain: " + e.Error())
		return "", errors.New("unable to parse hostname")
	}

	return r.Hostname(), nil
}

func (a *Adapter) handleSessionID(ctx context.Context, core *pmrOut) (*pmrOut, error) {
	hashedSessionID, e := a.encryptSessionID(ctx, core)
	if e != nil {
		a.monitor.LogHttpError(ctx, "api -> encryptSessionID: "+e.Error())
		return nil, e
	}

	core.SessionID = hashedSessionID
	return core, nil
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

func (a *Adapter) proxyValidatedEmailResponse(ctx context.Context, data *validateEmailResponse) (*validatedEmailIn, context.Context) {
	var workflowLabel = memberTypes.ContextPreRegistrationWorkflowLabel("proxyValidatedEmailResponse")

	ctx = context.WithValue(ctx, workflowLabel, "called")

	coreData := memberTypes.VerifiedEmailIn{
		Email:             data.GetEmail(),
		Autocorrect:       data.GetAutocorrect(),
		Deliverability:    data.GetDeliverability(),
		QualityScore:      data.GetQualityScore(),
		IsValidFormat:     data.GetIsValidFormat(),
		IsFreeEmail:       data.GetIsFreeEmail(),
		IsDisposableEmail: data.GetIsDisposableEmail(),
		IsRoleEmail:       data.GetIsRoleEmail(),
		IsCatchallEmail:   data.GetIsCatchallEmail(),
		IsMxFound:         data.GetIsMxFound(),
		IsSmtpValid:       data.GetIsSmtpValid(),
	}

	return &coreData, ctx
}

func (a *Adapter) PreRegisterPrimaryMember(ctx context.Context, data *pmrStart, ch chan *pmrOut, ech chan error) {
	validatedEmailResponse, e := a.driven.ValidateEmailAddress(ctx, data, a.gRPC.EmailNotificationclient)
	if e != nil {
		ech <- e
		return
	}

	drivenProxyData, ctx := a.proxyValidatedEmailResponse(ctx, validatedEmailResponse)

	core, ctx, e := a.core.PreRegister(ctx, *drivenProxyData)
	if e != nil {
		ech <- errors.Wrap(e, "PreRegisterPrimaryMember -> core.PreRegister")
		return
	}

	core, e = a.handleSessionID(ctx, core)
	if e != nil {
		ech <- errors.Wrap(e, "session id hashing failed")
		return
	}

	if core.RegistrationPending {
		e = a.sendPreRegistrationEmail(ctx, core.SessionID, []string{core.MemberID})
		if e != nil {
			a.monitor.LogGenericError(e.Error())
			ech <- errors.New("unable to process verification email")
			return
		}

		ch <- core
		return
	}

	// create and register account asynchronously
	ch <- core
	return
}
