package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	memberT "github.com/codeclout/AccountEd/members/member-types"
	memberCore "github.com/codeclout/AccountEd/members/ports/core"
	"github.com/codeclout/AccountEd/members/ports/framework/driven"
	notificationEmailv1 "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	sessionMembersv1 "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type cc = context.Context
type config = map[string]interface{}
type corePort = memberCore.MemberCorePort

type drivenPort = driven.MemberDrivenPort
type gRPCclients = protocol.AdapterServiceClients

type pmi = memberT.PrimaryMemberStartRegisterIn
type pmo = memberT.ValidatedEmailResonse

type MemberErrorOut = memberT.MemberErrorOut
type PMConfirmOut = memberT.PrimaryMemberConfirmationOut
type ValidatedEmailResonse = memberT.ValidatedEmailResonse

type Adapter struct {
	config             config
	contextAPILabel    memberT.ContextAPILabel
	contextDrivenLabel memberT.ContextDrivenLabel
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

func (a *Adapter) getEmailValidationToken(ctx cc, core *ValidatedEmailResonse) (*ValidatedEmailResonse, *MemberErrorOut) {
	if a.gRPC == nil || a.gRPC.MemberSessionclient == nil {
		return nil, &MemberErrorOut{Error: true, Msg: errors.New("nil gRPC or MemberSessionClient").Error()}
	}

	if core == nil {
		return nil, &MemberErrorOut{Error: true, Msg: "input parameter is nil"}
	}

	payload := sessionMembersv1.GenerateTokenRequest{
		HasAutoCorrect: len(core.AutoCorrect) > 0,
		MemberId:       core.MemberID,
		TokenId:        core.TokenID,
	}

	client := *a.gRPC.MemberSessionclient

	t, e := client.GenerateMemberToken(ctx, &payload)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, &MemberErrorOut{Error: true, Msg: "failed to encrypt session ID"}
	}

	core.Token = t.GetToken()

	return core, nil
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

func (a *Adapter) sendPrimaryMemberEmail(ctx context.Context, hashedSessionID string, toAddress []string) *MemberErrorOut {
	emailclient := *a.gRPC.EmailNotificationclient

	fqdn, e := a.getEmailDomain()
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return &memberT.MemberErrorOut{Error: true, Msg: "failed to get domain for email"}
	}

	uri, e := a.getDomain()
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return &memberT.MemberErrorOut{Error: true, Msg: "failed to get domain"}
	}

	reqData := notificationEmailv1.NoReplyEmailNotificationRequest{
		AwsCredentials: a.getAWSCredentialBytes(ctx),
		Domain:         uri,
		FromAddress:    fmt.Sprintf("no-reply@%s", fqdn),
		SessionId:      hashedSessionID,
		ToAddress:      toAddress,
	}

	_, e = emailclient.SendPreRegistrationEmail(ctx, &reqData)
	if e != nil {
		a.monitor.LogGenericError(fmt.Sprintf("pre-registration email failed to process -> %s", e.Error()))
		return &memberT.MemberErrorOut{Error: true, Msg: fmt.Sprintf("%s pre registration email failed", uri)}
	}

	return nil
}

func (a *Adapter) GetEmailVerificationConfirmation(ctx cc, token string, ch chan *PMConfirmOut, ech chan *MemberErrorOut) {
	pmConfirmOut, e := a.driven.ConfirmEmailAddressValidation(ctx, a.gRPC.MemberSessionclient, token)

	if e != nil {
		ech <- e
		return
	}

	ch <- pmConfirmOut
}

func (a *Adapter) VerifyPrimaryMemberEmail(ctx cc, data *pmi, ch chan *pmo, ech chan *memberT.MemberErrorOut) {
	emailValidationResp, e := a.driven.ValidateEmailAddress(ctx, data, a.gRPC.EmailNotificationclient)
	if e != nil {
		ech <- e
		return
	}

	core, e := a.core.ProcessEmailValidationResponse(ctx, emailValidationResp)
	if e != nil {
		ech <- e
		return
	}

	core, e = a.getEmailValidationToken(ctx, core)
	if e != nil {
		ech <- e
		return
	}

	if core.RegistrationPending {
		e = a.sendPrimaryMemberEmail(ctx, core.Token, []string{core.MemberID})
		if e != nil {
			ech <- e
			return
		}

		ch <- core
		return
	}

	// create and register account asynchronously
	ch <- core
}
