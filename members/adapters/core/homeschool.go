package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type Adapter struct {
	isRegistrable
	config  map[string]interface{}
	monitor monitoring.Adapter
}

type isRegistrable bool

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

// useAutoCorrect determines if Autocorrect is applicable for the input autocorrectString. It returns true if the length of autocorrectString
// is greater than 0, otherwise it returns false.
func (i *isRegistrable) useAutoCorrect(autocorrectString string) bool {
	return len(autocorrectString) > 0
}

func (i *isRegistrable) setRegistrationPending(delivery, qscore string, hasMX, isDisposable, isRole bool) bool {
	quality, _ := strconv.ParseFloat(qscore, 32)

	return quality < 0.80 || delivery != "DELIVERABLE" || isDisposable || isRole || !hasMX
}

func (a *Adapter) PreRegister(ctx context.Context, in memberTypes.VerifiedEmailIn) (*memberTypes.PrimaryMemberStartRegisterOut, context.Context, error) {
	var memberID, username string
	var workflowLabel = memberTypes.ContextPreRegistrationWorkflowLabel("core -> PreRegister")

	ctx = context.WithValue(ctx, workflowLabel, "called")

	sessionID, e := uuid.NewRandom()
	if e != nil {
		a.monitor.LogGenericError(fmt.Sprintf("error setting Session ID -> %s", e.Error()))
		return nil, ctx, errors.New("error setting session id")
	}

	if a.useAutoCorrect(in.Autocorrect) {
		memberID = ""

		x, e := memberTypes.ValidateEmail(&in.Autocorrect)
		if e != nil {
			a.monitor.LogGenericError(fmt.Sprintf("error validating auto corrected email address -> %s", e.Error()))
			return nil, ctx, errors.New("error validating auto corrected email address")
		}

		suggestedMemberID, ok := x.Load().(string)
		if !ok {
			a.monitor.LogGenericError(fmt.Sprintf("error loading auto corrected email address -> %s", e.Error()))
			return nil, ctx, errors.New("error loading auto corrected email address")
		}

		username = suggestedMemberID
	} else {
		memberID = in.Email
	}

	out := memberTypes.PrimaryMemberStartRegisterOut{
		AutoCorrect:         username,
		MemberID:            memberID,
		RegistrationPending: a.setRegistrationPending(in.Deliverability, in.QualityScore, in.IsMxFound.GetValue(), in.IsDisposableEmail.GetValue(), in.IsRoleEmail.GetValue()),
		SessionID:           sessionID.String(),
		UsernamePending:     a.useAutoCorrect(in.Autocorrect),
	}

	return &out, ctx, nil
}
