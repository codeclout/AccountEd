package core

import (
  "context"

  memberTypes "github.com/codeclout/AccountEd/members/member-types"
)

type Member interface {
  MemberGroupById(ctx context.Context, id string) *memberTypes.MemberGroup
  MemberTypeById(ctx context.Context, id string) *memberTypes.MemberType
  NewMemberSession(ctx context.Context) *memberTypes.MemberSession
  RefreshMemberSession(ctx context.Context, session *memberTypes.MemberSession) *memberTypes.MemberSession
  RevokeMemberSession(ctx context.Context, session *memberTypes.MemberSession) bool
}
