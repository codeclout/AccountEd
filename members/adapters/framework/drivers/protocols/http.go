package protocols

import (
  "golang.org/x/exp/slog"
)

type Adapter struct {
  log *slog.Logger
}

func NewAdapter() *Adapter {
  return &Adapter{}
}

func (a *Adapter) Run() {
  
}
