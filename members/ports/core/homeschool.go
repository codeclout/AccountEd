package core

import (
  "context"

  memberTypes "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolCore interface {
  ParentGuardiansByAccountId(ctx context.Context, id string) ([]memberTypes.ParentGuardian, error)
  ParentGuardianById(ctx context.Context, id string) (memberTypes.ParentGuardian, error)
  ParentGuardianByUsername(ctx context.Context, username string) (memberTypes.ParentGuardian, error)
  Register(ctx context.Context, in memberTypes.HomeSchoolRegisterIn) (memberTypes.HomeSchoolRegisterOut, error)
  StudentByMemberId(ctx context.Context, id string) (memberTypes.Student, error)
  StudentByPin(ctx context.Context, pin string, principal memberTypes.Member) (memberTypes.Student, error)
  StudentsByAccountId(ctx context.Context, id string) ([]memberTypes.Student, error)
}
