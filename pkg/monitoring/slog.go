package monitoring

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
)

type Adapter struct {
	Logger    *slog.Logger
}

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

	return &Adapter{Logger: logger}
}

func getTimeStamp() time.Time {
	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()

	return t
}

func (a *Adapter) GetTimeStamp() time.Time {
	return getTimeStamp()
}

func (a *Adapter) Initialize() {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	sl := <-s

	a.Logger.Warn(fmt.Sprintf("Signal %s - shutting down ", sl))
}

func (a *Adapter) HttpMiddlewareLogger(msg ...interface{}) {
	a.Logger.Info("request", msg...)
}
