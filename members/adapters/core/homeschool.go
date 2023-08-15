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
	isRegisterable
	config map[string]interface{}
	monitor monitoring.Adapter
}

type isRegisterable bool

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config: config,
		monitor: monitor,
	}
}

// useAutoCorrect determines if Autocorrect is applicable for the input autocorrectString. It returns true if the length of autocorrectString
// is greater than 0, otherwise it returns false.
func (i *isRegisterable) useAutoCorrect(autocorrectString string) bool {
	return len(autocorrectString) > 0
}

// setRegistrationPending checks deliverability, quality score, existence of MX record, disposable status, and role-based status to set registration pending
// for an email address. It returns false if the quality score is greater or equal to 0.80, delivery status is "DELIVERABLE", the email
// is not disposable or role-based, and has an MX record. Returns true otherwise.
func (i *isRegisterable) setRegistrationPending(delivery, qscore string, hasMX, isDisposable, isRole bool) bool {
	quality, _ := strconv.ParseFloat(qscore, 32)

	if quality >= 0.80 && delivery == "DELIVERABLE" && isDisposable == false && isRole == false && hasMX {
		return false
	}

	return true
}

func (a *Adapter) PreRegister(ctx context.Context) (*memberTypes.PrimaryMemberStartRegisterOut, error) {
	var memberID string
	var username string

	in, ok := ctx.Value(memberTypes.ContextAPILabel("api_input")).(memberTypes.EmailValidationIn)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid core parameters -> %v", ctx))
	}

	sessionID, e := uuid.NewRandom()
	if e != nil {
		return nil, errors.New("unable to set sessionID")
	}

	if a.useAutoCorrect(in.Autocorrect) {
		memberID = ""

		x, _ := memberTypes.ValidateEmail(&in.Autocorrect)
		suggestedMemberID := x.Load().(string)
		username = suggestedMemberID
	} else {
		memberID = in.Email
	}

	return &memberTypes.PrimaryMemberStartRegisterOut{
		AutoCorrect: username,
		MemberID:    memberID,
		RegistrationPending: a.setRegistrationPending(in.Deliverability, in.QualityScore, in.IsMxFound.GetValue(), in.IsDisposableEmail.GetValue(), in.IsRoleEmail.GetValue()),
		SessionID:           sessionID.String(),
		UsernamePending:     a.useAutoCorrect(in.Autocorrect),
	}, nil
}
