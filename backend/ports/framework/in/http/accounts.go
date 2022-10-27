package ports

import "github.com/gofiber/fiber/v2"

type UserAccountPort interface {
	HandleCreateAccountType(c *fiber.Ctx) error
}
