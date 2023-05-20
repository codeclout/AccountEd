package api

import (
  "context"

  "github.com/pkg/errors"
  "golang.org/x/exp/slog"

  memberTypes "github.com/codeclout/AccountEd/members/member-types"
  "github.com/codeclout/AccountEd/members/ports/core"
  "github.com/codeclout/AccountEd/pkg/monitoring"
)

type Adapter struct {
  core core.HomeschoolCore
  log  *slog.Logger
}

func NewAdapter(core core.HomeschoolCore, monitor *monitoring.Adapter) *Adapter {
  return &Adapter{
    core: core,
    log:  monitor.Logger,
  }
}

func (a *Adapter) RegisterAccount(ctx context.Context, data *memberTypes.HomeSchoolRegisterIn, ch chan memberTypes.HomeSchoolRegisterOut, ech chan error) {
  out, e := a.core.Register(ctx, data)
  if e != nil {
    ech <- errors.Wrapf(e, "registerAccountAPI -> core.Register(%v)", data)
  }

  ch <- out
}
