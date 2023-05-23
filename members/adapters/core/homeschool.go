package core

import (
  "context"

  "golang.org/x/exp/slog"

  mt "github.com/codeclout/AccountEd/members/member-types"
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

func (a *Adapter) PreRegister(ctx context.Context, in mt.PrimaryMemberStartRegisterIn) (*mt.PrimaryMemberStartRegisterOut, error) {
  return nil, nil
}

