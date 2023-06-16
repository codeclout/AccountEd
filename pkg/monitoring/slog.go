package monitoring

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"golang.org/x/exp/slog"
)

type Adapter struct {
	Logger *slog.Logger
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

	return &Adapter{Logger: logger}
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

// Initialize sets up a signal interrupt handler for the Adapter, capturing SIGINT and SIGTERM signals to gracefully shut down the Logger.
// Upon receiving any of these signals, the function logs a warning message that includes the received signal's value.
func (a *Adapter) Initialize(wg *sync.WaitGroup) {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	sl := <-s

	a.Logger.Warn(fmt.Sprintf("Signal %s - shutting down ", sl))
	wg.Done()
}

func (a *Adapter) HttpMiddlewareLogger(msg string, attr slog.Attr) {
	a.Logger.Info(msg, attr)
}
