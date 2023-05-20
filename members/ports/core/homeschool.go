package core

import (
  "context"

  membertypes "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolCore interface {
  //ParentGuardiansByAccountId(ctx context.Context, id string) ([]membertypes.ParentGuardian, error)
  //ParentGuardianById(ctx context.Context, id membertypes.MemberID) (membertypes.ParentGuardian, error)
  //ParentGuardianByUsername(ctx context.Context, username string) (membertypes.ParentGuardian, error)
  Register(ctx context.Context, in *membertypes.HomeSchoolRegisterIn) (membertypes.HomeSchoolRegisterOut, error)
  //StudentByMemberId(ctx context.Context, id membertypes.MemberID) (membertypes.Student, error)
  //StudentByPin(ctx context.Context, pin string, principal membertypes.MemberID) (membertypes.Student, error)
  //StudentsByAccountId(ctx context.Context, id string) ([]membertypes.Student, error)
}
