package api

import (
  "context"

  "github.com/pkg/errors"
  "golang.org/x/exp/slog"

  mt "github.com/codeclout/AccountEd/members/member-types"
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

func (a *Adapter) PreRegisterPrimaryMember(ctx context.Context, data *mt.PrimaryMemberStartRegisterIn, ch chan *mt.PrimaryMemberStartRegisterOut, ech chan error) {
  out, e := a.core.PreRegister(ctx, *data)
  if e != nil {
    ech <- errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", data)
    ctx.Done()
  }

  ch <- out
}
