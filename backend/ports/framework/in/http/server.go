package ports

import "github.com/gofiber/fiber/v2"

type HTTPPort interface {
	Run()
	CreateAccountType(c *fiber.Ctx) error
}
