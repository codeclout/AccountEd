package core

import (
  "context"

  "golang.org/x/exp/slog"

  mt "github.com/codeclout/AccountEd/members/member-types"
)

type Adapter struct {
  log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
  return &Adapter{
    log: log,
  }
}

func (a *Adapter) PreRegister(ctx context.Context, in mt.PrimaryMemberStartRegisterIn) (*mt.PrimaryMemberStartRegisterOut, error) {
  return nil, nil
}

