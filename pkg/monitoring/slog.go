package monitoring

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	monitorTypes "github.com/codeclout/AccountEd/pkg/monitoring/monitoring-types"
)

type Adapter struct {
	LogLabelRequestID     monitorTypes.LogLabel
	LogLabelTransactionID monitorTypes.LogLabel
	XForwardedFor         monitorTypes.LogLabel
	Logger                *slog.Logger
}

// NewAdapter is a constructor function that returns a new Adapter instance. It initializes the logger based on the environment
func NewAdapter() *Adapter {
	var logger *slog.Logger

	o := slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := getTimeStamp()
				a.Value = slog.TimeValue(t)
			}
			return a
		}}

	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && strings.TrimSpace(env) == "prod" {
		o.Level = slog.LevelInfo
		h := slog.NewJSONHandler(os.Stdout, &o)

		logger = slog.New(h)
		slog.SetDefault(logger)

	} else {
		h := slog.NewJSONHandler(os.Stdout, &o)
		logger = slog.New(h)
		slog.SetDefault(logger)
	}

	return &Adapter{
		LogLabelRequestID:     "request_id",
		LogLabelTransactionID: "transaction_id",
		Logger:                logger,
		XForwardedFor:         "forwarded_ip",
	}
}

// getTimeStamp returns the current timestamp in UTC format, with nanosecond precision. This function is used
// internally by the monitoring package to generate time values for logging purposes. The returned time value represents the
// moment the function is called, and can be used with other functions in the time package for time manipulation. Note that this
// function does not have any input parameters, nor is it a method of any type.
func getTimeStamp() time.Time {
	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()

	return t
}

func (a *Adapter) GetTimeStamp() time.Time {
	return getTimeStamp()
}

func (a *Adapter) HttpMiddlewareLogger(msg string, attr slog.Attr) {
	a.Logger.Info(msg, attr)
}

func (a *Adapter) LogHttpError(ctx context.Context, msg string) {
	a.Logger.Error(msg,
		"request_id", ctx.Value(monitorTypes.LogLabel("request_id")),
		"transaction_id", fmt.Sprintf("%x", ctx.Value(monitorTypes.LogLabel("transaction_id"))))
}

func (a *Adapter) LogHttpInfo(ctx context.Context, msg string) {
	a.Logger.Info(msg,
		"request_id", ctx.Value(monitorTypes.LogLabel("request_id")),
		"transaction_id", fmt.Sprintf("%x", ctx.Value(monitorTypes.LogLabel("transaction_id"))))
}

func (a *Adapter) LogGenericError(msg string) {
	a.Logger.Error(msg)
}

func (a *Adapter) LogGrpcError(ctx context.Context, msg string) {
	a.Logger.Error(msg,
		"transaction_id", fmt.Sprintf("%x", ctx.Value(monitorTypes.LogLabel("transaction_id"))))
}

func (a *Adapter) LogGrpcInfo(ctx context.Context, msg string) {
	a.Logger.Info(msg,
		"transaction_id", fmt.Sprintf("%x", ctx.Value(monitorTypes.LogLabel("transaction_id"))))
}

func (a *Adapter) LogGenericInfo(msg string) {
	a.Logger.Info(msg)
}
