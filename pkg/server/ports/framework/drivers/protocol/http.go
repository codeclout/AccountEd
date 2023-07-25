package protocol

import (
	"github.com/gofiber/fiber/v2"
)

type ServerProtocolHttpPort interface {
	GetPort() (string, error)
	InitializeRoutes(routes []*fiber.App)
}
