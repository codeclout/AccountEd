package monitoring

import (
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
)

type Adapter struct {
	Logger    *slog.Logger
	WaitGroup *sync.WaitGroup
}

func NewAdapter() *Adapter {
	o := slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := GetTimeStamp()
				a.Value = slog.TimeValue(t)
			}
			return a
		}}

	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && strings.TrimSpace(env) == "prod" {
		o.Level = slog.LevelInfo
		h := o.NewJSONHandler(os.Stdout)

		logger := slog.New(h)
		return &Adapter{Logger: logger}
	}

	h := o.NewJSONHandler(os.Stdout)
	logger := slog.New(h)

	return &Adapter{Logger: logger}
}

func GetTimeStamp() time.Time {
	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()

	return t
}

func (a *Adapter) Initialize() {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	sl := <-s

	a.Logger.Warn("Signal %s - shutting down: ", sl)
	a.WaitGroup.Wait()

}

func (a *Adapter) HttpMiddlewareLogger(msg ...interface{}) {
	a.Logger.Info("request", msg...)
}
