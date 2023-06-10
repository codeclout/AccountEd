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

func (i *isRegisterable) useAutoCorrect(autocorrectString string) bool {
	return len(autocorrectString) > 0
}

func (i *isRegisterable) setRegistrationPending(delivery, qscore string, hasMX, isDisposable, isRole bool) bool {
	quality, _ := strconv.ParseFloat(qscore, 32)

	if quality >= 0.80 && delivery == "DELIVERABLE" && isDisposable == false && isRole == false && hasMX {
		return false
	}

	return true
}

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
