package core

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	emailv1 "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type ValidateEmailAddressResponse = emailv1.ValidateEmailAddressResponse

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) NewMemberSession(ctx context.Context) *memberTypes.MemberSession {
	return nil
}

func (a *Adapter) MemberGroupByID(ctx context.Context, id string) *memberTypes.MemberGroup {
	return nil
}

func (a *Adapter) MemberTypeByID(ctx context.Context, id string) *memberTypes.MemberType {
	return nil
}

func (a *Adapter) ProcessEmailValidationResponse(ctx context.Context, in *ValidateEmailAddressResponse) (*memberTypes.ValidatedEmailResonse, *memberTypes.MemberErrorOut) {
	tokenID, e := uuid.NewRandom()
	if e != nil {
		const msg = "error setting Session ID"
		a.monitor.LogHttpError(ctx, fmt.Sprintf(msg+": %s", e.Error()))
		return nil, &memberTypes.MemberErrorOut{Error: true, Msg: msg}
	}

	out := memberTypes.ValidatedEmailResonse{
		AutoCorrect:         in.GetAutoCorrect(),
		MemberID:            in.GetMemberId(),
		RegistrationPending: in.GetShouldConfirmAddress(),
		TokenID:             tokenID.String(),
		UsernamePending:     in.GetMemberIdPending(),
	}

	return &out, nil
}

func (a *Adapter) RefreshMemberSession(ctx context.Context, session *memberTypes.MemberSession) *memberTypes.MemberSession {
	return nil
}

func (a *Adapter) RevokeMemberSession(ctx context.Context, session *memberTypes.MemberSession) bool {
	return false
}
