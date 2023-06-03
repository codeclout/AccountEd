package protocols

import (
  "github.com/gofiber/fiber/v2"

  httpMiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
)

type httpProtocol interface {
  Run(middlewareLogger func(settings ...httpMiddleware.Config) fiber.Handler)
}
