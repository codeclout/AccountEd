package core

import (
	"context"
	"strconv"

	"golang.org/x/exp/slog"

	"github.com/google/uuid"

	mt "github.com/codeclout/AccountEd/members/member-types"
)

type Adapter struct {
	isRegisterable
	log *slog.Logger
}

type isRegisterable bool

func NewAdapter(log *slog.Logger) *Adapter {
	return &Adapter{
		log: log,
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

// PreRegister performs the pre-registration process for an email address, determining if it is deliverable, has a valid MX record,
// is not disposable, or a role-based email. It sets the registration state as pending based on these criteria and generates a session ID.
// Returns a PrimaryMemberStartRegisterOut object with registration pending status, session ID, username, and a bool indicating if username is
// pending autocorrection. An error is returned if any issue occurs during this process.
func (a *Adapter) PreRegister(ctx context.Context, in mt.EmailValidationIn) (*mt.PrimaryMemberStartRegisterOut, error) {
	var username *string
	sessionID, _ := uuid.NewRandom()

	if a.useAutoCorrect(in.Autocorrect) {
		x, _ := mt.ValidateEmail(&in.Autocorrect)
		v := x.Load().(string)
		username = &v
	} else {
		username = &in.Email
	}

	return &mt.PrimaryMemberStartRegisterOut{
		RegistrationPending: a.setRegistrationPending(in.Deliverability, in.QualityScore, in.IsMxFound.GetValue(), in.IsDisposableEmail.GetValue(), in.IsRoleEmail.GetValue()),
		SessionID:           sessionID,
		Username:            username,
		UsernamePending:     a.useAutoCorrect(in.Autocorrect),
	}, nil
}
