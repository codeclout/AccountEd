package logger

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Adapter struct {
	Logger *zap.SugaredLogger
}

func NewAdapter() *Adapter {
	var logger *zap.Logger

	z := []byte(`{
		"encoding": "json",
		"errorOutputPaths": ["stderr"],
		"outputPaths": ["stdout"],
		"encoderConfig": { "levelKey": "level", "messageKey": "msg", "timeKey": "timestamp" }
	}`)

	if en, ok := os.LookupEnv("ENVIRONMENT"); ok && strings.TrimSpace(en) == "prod" {
		c := zap.NewProductionConfig()

		if e := json.Unmarshal(z, &c); e != nil {
			log.Fatalf("%v", e)
		}

		c.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		c.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
		c.EncoderConfig.StacktraceKey = ""

		logger, _ = c.Build()

		return &Adapter{
			Logger: logger.Sugar(),
		}
	}

	c := zap.NewDevelopmentConfig()

	if e := json.Unmarshal(z, &c); e != nil {
		log.Fatalf("%v", e)
	}

	c.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	c.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	c.EncoderConfig.StacktraceKey = ""

	logger, _ = c.Build()

	return &Adapter{
		Logger: logger.Sugar(),
	}
}

func (a *Adapter) Initialize() {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	sl := <-s

	a.Logger.Warn("Signal %s - shutting down: ", sl)
}

func (a *Adapter) Log(level string, msg string) {
	if ok, e := regexp.MatchString(`^debug$`, level); ok && e == nil {
		a.Logger.Debug(msg)
	}

	if ok, e := regexp.MatchString(`^error$`, level); ok && e == nil {
		a.Logger.Error(msg)
	}

	if ok, e := regexp.MatchString(`^fatal$`, level); ok && e == nil {
		a.Logger.Fatal(msg)
	}

	if ok, e := regexp.MatchString(`^info$`, level); ok && e == nil {
		a.Logger.Info(msg)
	}

	if ok, e := regexp.MatchString(`^warn$`, level); ok && e == nil {
		a.Logger.Warn(msg)
	}
}

func (a *Adapter) HttpMiddlewareLogger(msg ...interface{}) {
	a.Logger.Infow("request", msg...)
}

func (a *Adapter) Sync() {
	_ = a.Logger.Sync()
}
