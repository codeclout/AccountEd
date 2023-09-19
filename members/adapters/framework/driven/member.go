package driven

import (
	"context"
	"fmt"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	membersessionv1 "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type cc = context.Context
type ec = emailv1.EmailNotificationServiceClient
type MErrorOut = memberTypes.MemberErrorOut
type pm = memberTypes.PrimaryMemberStartRegisterIn
type PMConfirmOut = memberTypes.PrimaryMemberConfirmationOut
type sc = membersessionv1.MemberSessionClient

type ValidateEmailAddressResponse = emailv1.ValidateEmailAddressResponse

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{monitor: monitor}
}

func (a *Adapter) ValidateEmailAddress(ctx cc, data *pm, emailClient *ec) (*ValidateEmailAddressResponse, *MErrorOut) {
	if v := *emailClient; v == nil {
		const msg = "nil notifications gRPC client"
		a.monitor.LogGenericError(msg)
		return nil, &memberTypes.MemberErrorOut{Error: true, Msg: msg}
	}

	if x := data; x == nil {
		const msg = "nil primary member email address"
		a.monitor.LogGenericError(msg)
		return nil, &memberTypes.MemberErrorOut{Error: true, Msg: msg}
	}

	client := *emailClient
	response, e := client.ValidateEmailAddress(ctx, &emailv1.ValidateEmailAddressRequest{Address: *data.Username})

	if e != nil {
		const msg = "error validating primary member"
		a.monitor.LogGrpcError(ctx, fmt.Sprintf(msg+": %v => %s", *data, e.Error()))
		return nil, &memberTypes.MemberErrorOut{Error: true, Msg: msg}
	}

	return response, nil
}

func (a *Adapter) ConfirmEmailAddressValidation(ctx cc, sessionClient *sc, token string) (*PMConfirmOut, *MErrorOut) {
	var s string

	if v := *sessionClient; v == nil {
		const msg = "nil session gRPC client"
		a.monitor.LogGenericError(msg)
		return nil, &memberTypes.MemberErrorOut{Error: true, Msg: msg}
	}

	if token == (s) {
		const msg = "empty encrypted string"
		a.monitor.LogGenericError(msg)
		return nil, &memberTypes.MemberErrorOut{Error: true, Msg: msg}
	}

	client := *sessionClient

	isValid, e := client.ValidateMemberToken(ctx, &membersessionv1.ValidateTokenRequest{Token: token})
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, &MErrorOut{Error: true, Msg: e.Error()}
	}

	return &PMConfirmOut{IsPrimaryMemberConfirmed: isValid.GetIsValidToken()}, nil
}
