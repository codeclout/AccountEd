package drivers

import (
	"context"

	"golang.org/x/exp/slog"
)

type SlogLoggerDriverPort interface {
	HttpMiddlewareLogger(msg string, attr slog.Attr)
	LogHttpError(ctx context.Context, msg string)
	LogGenericError(msg string)
	LogGrpcError(ctx context.Context, msg string)
	LogGenericInfo(msg string)
}
