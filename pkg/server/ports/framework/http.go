package framework

import "github.com/gofiber/fiber/v2"

type ProtocolHTTP interface {
  Initialize(routePrefix string, config *fiber.Config)
}
