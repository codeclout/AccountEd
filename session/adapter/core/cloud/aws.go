package cloud

import (
	"context"

	"golang.org/x/exp/slog"
)

type Adapter struct {
	log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
	return &Adapter{
		log: log,
	}
}

func (a *Adapter) GetServiceIdMetadata(ctx context.Context) {
	
}
