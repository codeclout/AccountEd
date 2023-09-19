package api

import (
	"context"

	memberT "github.com/codeclout/AccountEd/members/member-types"
)

type cc = context.Context
type pmi = memberT.PrimaryMemberStartRegisterIn
type pmo = memberT.ValidatedEmailResonse

type MemberAPI interface {
	GetEmailVerificationConfirmation(ctx cc, token string, ch chan *memberT.PrimaryMemberConfirmationOut, ech chan *memberT.MemberErrorOut)
	VerifyPrimaryMemberEmail(ctx cc, data *pmi, ch chan *pmo, ech chan *memberT.MemberErrorOut)
}
