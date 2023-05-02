package api

import (
  "context"

  memberTypes "github.com/codeclout/AccountEd/members/member-types"
)

type Adapter struct{}

func NewAdapter() *Adapter {
  return &Adapter{}
}

func (a *Adapter) GetParentGuardiansByAccountId(ctx context.Context, id string, in chan<- []memberTypes.ParentGuardian) {
}

func (a *Adapter) GetParentGuardianById(ctx context.Context, id string, in chan<- memberTypes.ParentGuardian) {
}

func (a *Adapter) GetParentGuardianByUsername(ctx context.Context, username string, in chan<- memberTypes.ParentGuardian) {
}

func (a *Adapter) RegisterAccount(ctx context.Context, data memberTypes.HomeSchoolRegisterIn, in chan<- memberTypes.HomeSchoolRegisterOut) {
}

func (a *Adapter) GetStudentByMemberId(ctx context.Context, id string, in chan<- memberTypes.Student) {
}

func (a *Adapter) GetStudentByPin(ctx context.Context, pin string, principal memberTypes.Member, in chan<- memberTypes.Student) {
}

func (a *Adapter) GetStudentsByAccountId(ctx context.Context, id string, in chan<- []memberTypes.Student) {
}
