package core

import (
  "context"
  "errors"

  "github.com/codeclout/AccountEd/members/member-types"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
  return &Adapter{}
}

func (a *Adapter) Register(ctx context.Context, in memberTypes.HomeSchoolRegisterIn) (memberTypes.HomeSchoolRegisterOut, error) {
  return memberTypes.HomeSchoolRegisterOut{}, errors.New("not implemented")
}

func (a *Adapter) ParentGuardiansByAccountId(ctx context.Context, id string) ([]memberTypes.ParentGuardian, error) {
  return nil, errors.New("not Implemented")
}

func (a *Adapter) ParentGuardianById(ctx context.Context, id string) (memberTypes.ParentGuardian, error) {
  return memberTypes.ParentGuardian{}, errors.New("not Implemented")
}

func (a *Adapter) ParentGuardianByUsername(ctx context.Context, username string) (memberTypes.ParentGuardian, error) {
  return memberTypes.ParentGuardian{}, errors.New("not implemented")
}

func (a *Adapter) StudentByMemberId(ctx context.Context, id string) (memberTypes.Student, error) {
  return memberTypes.Student{}, errors.New("not implemented")
}

func (a *Adapter) StudentByPin(ctx context.Context, pin string, principal memberTypes.Member) (memberTypes.Student, error) {
  return memberTypes.Student{}, errors.New(("not implemented"))
}

func (a *Adapter) StudentsByAccountId(ctx context.Context, id string) ([]memberTypes.Student, error) {
  return nil, errors.New(("not implemented"))
}
