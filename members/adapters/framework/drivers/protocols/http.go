package protocols

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"

	httpMiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
)

type Adapter struct {
	app              *fiber.App
	log              *slog.Logger
	middlewareLogger func(settings ...httpMiddleware.Config) fiber.Handler
}

func NewAdapter(log *slog.Logger, app *fiber.App, middlewareLogger func(settings ...httpMiddleware.Config) fiber.Handler) *Adapter {
	return &Adapter{
		app:              app,
		log:              log,
		middlewareLogger: middlewareLogger,
	}
}

func (a *Adapter) InitializeClient(port string) {
	a.log.Info("starting server")
	log.Fatal(a.app.Listen(port))
}
