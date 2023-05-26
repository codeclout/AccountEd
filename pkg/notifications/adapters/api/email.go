package api

import (
  "context"

  "github.com/pkg/errors"
  "golang.org/x/exp/slog"

  pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/core"
)

type Adapter struct {
  core core.EmailCorePort
  log  *slog.Logger
}

func NewAdapter(log *slog.Logger, core core.EmailCorePort) *Adapter {
  return &Adapter{
    core: core,
    log:  log,
  }
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, errorch chan error) {
  var reason []*pb.Reason

  ctx = context.WithValue(ctx, "address", address)
  validated, e := a.core.ProcessEmailValidation(ctx)

  if e != nil {
    x := errors.Wrapf(e, "api-ValidateEmailAddress -> core.ProcessEmailValidation(%v)", ctx)
    errorch <- x
  }

  for _, r := range validated.Reason {
    x := &pb.Reason{Reason: r}
    reason = append(reason, x)
  }

  out := &pb.ValidateEmailAddressResponse{
    SessionId:     nil,
    Address:       validated.Address,
    IsDisposable:  validated.IsDisposable,
    IsRoleAddress: validated.IsRoleAddress,
    Reason:        reason,
    Result:        validated.Result,
    Risk:          validated.Risk,
  }

  ch <- out
}
