package api

import (
  "context"

  memberTypes "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolAPI interface {
  //GetParentGuardiansByAccountId(ctx context.Context, id string, in chan<- []memberTypes.ParentGuardian)
  //GetParentGuardianById(ctx context.Context, id string, in chan<- memberTypes.ParentGuardian)
  //GetParentGuardianByUsername(ctx context.Context, username string, in chan<- memberTypes.ParentGuardian)
  RegisterAccount(ctx context.Context, data *memberTypes.HomeSchoolRegisterIn, ch chan memberTypes.HomeSchoolRegisterOut, ech chan error)
  //GetStudentByMemberId(ctx context.Context, id string, in chan<- memberTypes.Student)
  //GetStudentByPin(ctx context.Context, pin string, principal memberTypes.Member, in chan<- memberTypes.Student)
  //GetStudentsByAccountId(ctx context.Context, id string, in chan<- []memberTypes.Student)
}
