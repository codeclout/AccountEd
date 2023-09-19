package drivers

import "github.com/gofiber/fiber/v2"

type HomeschoolDriverPort interface {
	InitializeHomeschoolAPI() []*fiber.App
}
