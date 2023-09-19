package core

import (
	"context"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	emailv1 "github.com/codeclout/AccountEd/notifications/gen/email/v1"
)

type cc = context.Context

type MemberCorePort interface {
	MemberCore
	MemberCoreActions
}

type MemberCore interface {
	ProcessEmailValidationResponse(ctx cc, in *emailv1.ValidateEmailAddressResponse) (*memberTypes.ValidatedEmailResonse, *memberTypes.MemberErrorOut)
}

type MemberCoreActions interface {
	MemberGroupByID(ctx context.Context, id string) *memberTypes.MemberGroup
	MemberTypeByID(ctx context.Context, id string) *memberTypes.MemberType
	NewMemberSession(ctx context.Context) *memberTypes.MemberSession
	RefreshMemberSession(ctx context.Context, session *memberTypes.MemberSession) *memberTypes.MemberSession
	RevokeMemberSession(ctx context.Context, session *memberTypes.MemberSession) bool
}
