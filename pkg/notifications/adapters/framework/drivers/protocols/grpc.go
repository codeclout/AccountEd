package protocols

import "golang.org/x/exp/slog"

type Adapter struct {
  log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
  return &Adapter{log: log}
}

func (a *Adapter) Run() {}
