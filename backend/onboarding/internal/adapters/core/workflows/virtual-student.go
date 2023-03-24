package workflows

import "context"

type logger func(level, msg string)

type Adapter struct {
  log logger
}

func NewAdapter(l logger) *Adapter {
  return &Adapter{
    log: l,
  }
}

func PublicSchoolsByZone(ctx context.Context) {}

func PublicSchoolById(ctx context.Context) {}
