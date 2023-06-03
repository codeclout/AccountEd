package protocols

import (
  "log"

  "github.com/gofiber/fiber/v2"
  "golang.org/x/exp/slog"

  httpMiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
)

type Adapter struct {
  app *fiber.App
  log *slog.Logger
}

func NewAdapter(log *slog.Logger, app *fiber.App) *Adapter {
  return &Adapter{
    app: app,
    log: log,
  }
}

func (a *Adapter) Run(middlewareLogger func(settings ...httpMiddleware.Config) fiber.Handler) {
  a.log.Info("starting server")

  a.app.Use(middlewareLogger(httpMiddleware.Config{
    Log:        middlewareLogger,
    ShouldSkip: nil,
  }))
  log.Fatal(a.app.Listen(":8088"))
}
