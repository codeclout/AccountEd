package core

import (
  "context"
  "errors"

  "golang.org/x/exp/slog"

  "github.com/codeclout/AccountEd/members/member-types"
  "github.com/codeclout/AccountEd/pkg/monitoring"
)

type Adapter struct {
  log *slog.Logger
}

func NewAdapter(monitor *monitoring.Adapter) *Adapter {
  return &Adapter{
    log: monitor.Logger,
  }
}

func (a *Adapter) Register(ctx context.Context, in *membertypes.HomeSchoolRegisterIn) (membertypes.HomeSchoolRegisterOut, error) {
  return membertypes.HomeSchoolRegisterOut{}, errors.New("not implemented")
}

func (a *Adapter) ParentGuardiansByAccountId(ctx context.Context, id string) ([]membertypes.ParentGuardian, error) {
  return nil, errors.New("not Implemented")
}

func (a *Adapter) ParentGuardianById(ctx context.Context, id string) (membertypes.ParentGuardian, error) {
  return membertypes.ParentGuardian{}, errors.New("not Implemented")
}

func (a *Adapter) ParentGuardianByUsername(ctx context.Context, username string) (membertypes.ParentGuardian, error) {
  return membertypes.ParentGuardian{}, errors.New("not implemented")
}

func (a *Adapter) StudentByMemberId(ctx context.Context, id string) (membertypes.Student, error) {
  return membertypes.Student{}, errors.New("not implemented")
}

func (a *Adapter) StudentByPin(ctx context.Context, pin string, principal membertypes.Member) (membertypes.Student, error) {
  return membertypes.Student{}, errors.New(("not implemented"))
}

func (a *Adapter) StudentsByAccountId(ctx context.Context, id string) ([]membertypes.Student, error) {
  return nil, errors.New(("not implemented"))
}
