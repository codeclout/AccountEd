package api

import (
  "context"

  "golang.org/x/exp/slog"

  pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/core"
)

type Adapter struct {
  core core.NotificationEmailCore
  log  *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
  return &Adapter{log: log}
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, errorch chan error) {
  a.core.ProcessEmailValidation(ctx)
}
