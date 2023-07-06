package protocols

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"

	httpMiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
)

type Adapter struct {
	app              *fiber.App
	config           map[string]interface{}
	log              *slog.Logger
	middlewareLogger func(settings ...httpMiddleware.Config) fiber.Handler
}

func NewAdapter(config map[string]interface{}, app *fiber.App, middlewareLogger func(settings ...httpMiddleware.Config) fiber.Handler, log *slog.Logger) *Adapter {
	return &Adapter{
		app:              app,
		config:           config,
		log:              log,
		middlewareLogger: middlewareLogger,
	}
}

func (a *Adapter) InitializeNotificationsClient(port string) {
	a.log.Info("starting server")
	log.Fatal(a.app.Listen(port))
}
