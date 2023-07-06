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
	"golang.org/x/exp/slog"

	httpMiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
)

type middlewareLogger func(msg string, attr slog.Attr)

type Adapter struct {
	HTTP                 *fiber.App
	WaitGroup            *sync.WaitGroup
	applicationName      string
	config               map[string]interface{}
	isApplicationGetOnly bool
	log                  *slog.Logger
	middlewareLogger     middlewareLogger
	routePrefix          string
}

// getPort retrieves the port number from the "PORT" environment variable and validates the value. If the value exists, fulfills a
// minimum length of 4, and falls within the valid port range (1024 - 65535), getPort returns a string formatted with the port number.
// If the value is not present or invalid, it returns the default port string ":8088".
func getPort() string {
	p, ok := os.LookupEnv("PORT")
	n, _ := strconv.Atoi(p)

	if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
		return fmt.Sprintf(":%d", n)
	}

	return ":8088"
}

// isProd checks if the current environment is a production environment. It does so by looking up the "ENVIRONMENT" environment
// variable and comparing its trimmed value to "prod". If it matches, the function returns true, otherwise it returns false.
func isProd() bool {
	if env, ok := os.LookupEnv("ENVIRONMENT"); ok && strings.TrimSpace(env) == "prod" {
		return true
	}

	return false
}

func NewAdapter(config map[string]interface{}, routePrefix, applicationName string, log *slog.Logger, mwl middlewareLogger, wg *sync.WaitGroup, isAppGetOnly bool) *Adapter {
	api := fiber.New()

	return &Adapter{
		HTTP:                 api,
		WaitGroup:            wg,
		applicationName:      applicationName,
		config:               config,
		isApplicationGetOnly: isAppGetOnly,
		log:                  log,
		middlewareLogger:     mwl,
		routePrefix:          routePrefix,
	}
}

// Initialize sets up a new fiber.App instance with the provided configurations and mounts the passed API endpoints to it.
// It returns the configured app and a string representation of the listening port number.
func (a *Adapter) Initialize(api []*fiber.App) (*fiber.App, string) {
	isAppGetOnly := func() int { // FixMe - write buffer size
		if a.isApplicationGetOnly {
			return 0
		}
		return 512
	}

	app := fiber.New(fiber.Config{
		AppName:                      a.applicationName,
		DisablePreParseMultipartForm: true,
		DisableStartupMessage:        isProd(),
		IdleTimeout:                  2 * time.Second,
		ReadTimeout:                  2 * time.Second,
		WriteBufferSize:              isAppGetOnly(),
	})

	for _, x := range api {
		a.log.Info("creating API routes")

		app.Use(httpMiddleware.NewLoggerMiddleware(httpMiddleware.Config{
			Log: a.middlewareLogger,
			ShouldSkip: func(c *fiber.Ctx) bool {
				return false
			},
		}))

		app.Mount(a.routePrefix, x)
	}

	return app, getPort()
}

// PostInit listens for SIGINT and SIGTERM signals, initiates shutdown, and stops the protocol listener for the given Fiber App.
func (a *Adapter) PostInit(wg *sync.WaitGroup) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	a.log.Warn("http client shutting down")

	wg.Done()
	os.Exit(0)
}

func (a *Adapter) StopProtocolListener(app *fiber.App) {
	a.WaitGroup.Wait()
	_ = app.Shutdown()
}
