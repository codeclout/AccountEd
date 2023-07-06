package member

import (
	"context"
	"golang.org/x/exp/slog"
)

type Adapter struct {
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, log *slog.Logger) *Adapter {
	return &Adapter{
		config: config,
		log:    log,
	}
}

func (a *Adapter) GetSessionIdKey(ctx context.Context, awsconfig []byte) ()
