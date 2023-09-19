package drivers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	memberT "github.com/codeclout/AccountEd/members/member-types"
)

type cc = context.Context
type MErrorOut = memberT.MemberErrorOut
type PrimaryMemberStartIn = memberT.PrimaryMemberStartRegisterIn
type PrimaryMemberStartOut = memberT.ValidatedEmailResonse

type MemberDriverPort interface {
	HandlePrimaryMemberValidity(ctx cc, in *PrimaryMemberStartIn) (*PrimaryMemberStartOut, *MErrorOut)
	HandleEmailVerification(ctx context.Context, token string) (*memberT.PrimaryMemberConfirmationOut, *memberT.MemberErrorOut)
	InitializeMemberAPI() []*fiber.App
}
