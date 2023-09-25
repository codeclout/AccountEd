package protocol

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/pkg/errors"

	httpMiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/server/server-types/protocols"
)

type Adapter struct {
	FrameworkDriver *fiber.App
	WaitGroup       *sync.WaitGroup
	config          map[string]interface{}
	metadata        protocols.ServerProtocolHttpMetadata
	monitor         monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, metadata protocols.ServerProtocolHttpMetadata, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	app := initialize(metadata)

	wg.Add(1)
	go postInit(monitor, wg, "http client shutting down")

	return &Adapter{
		FrameworkDriver: app,
		WaitGroup:       wg,
		config:          config,
		metadata:        metadata,
		monitor:         monitor,
	}
}

func isProd() bool {
	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && strings.TrimSpace(env) == "prod" {
		return true
	}

	return false
}

func initialize(metadata protocols.ServerProtocolHttpMetadata) *fiber.App {
	isAppGetOnly := func() int { // FixMe - write buffer size
		if metadata.UseOnlyGetRoutes {
			return 0
		}
		return 512
	}

	app := fiber.New(fiber.Config{
		AppName:                      metadata.ServerName,
		DisablePreParseMultipartForm: true,
		DisableStartupMessage:        isProd(),
		IdleTimeout:                  2 * time.Second,
		ReadTimeout:                  2 * time.Second,
		WriteBufferSize:              isAppGetOnly(),
	})

	return app
}

func postInit(monitor monitoring.Adapter, wg *sync.WaitGroup, warning string) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	monitor.Logger.Warn(warning)

	wg.Done()
	os.Exit(0)
}

func (a *Adapter) GetPort() (string, error) {
	port, ok := a.config["Port"].(string)
	if !ok {
		a.monitor.LogGenericError("port not configured")
		os.Exit(1)
	}

	n, _ := strconv.Atoi(port)

	if ok && len(strings.TrimSpace(port)) >= 4 && n >= 1024 && n <= 65535 {
		return fmt.Sprintf(":%d", n), nil
	}

	return "", errors.New("environment port out of bounds")
}

func (a *Adapter) InitializeRoutes(routes []*fiber.App) {
	a.FrameworkDriver.Use(httpMiddleware.NewLoggerMiddleware(httpMiddleware.Config{
		Log: a.monitor.HttpMiddlewareLogger,
		ShouldSkip: func(c *fiber.Ctx) bool {
			return false
		},
	}))

	a.FrameworkDriver.Use(requestid.New())

	for _, app := range routes {
		a.monitor.LogGenericInfo("creating API routes for: " + a.metadata.ServerName)
		a.FrameworkDriver.Mount("/v1/api/"+a.metadata.RoutePrefix, app)
	}
}
