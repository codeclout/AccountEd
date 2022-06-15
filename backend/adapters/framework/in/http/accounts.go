package http

import "github.com/gofiber/fiber/v2"

func (a Adapter) CreateAccountType(c *fiber.Ctx) error {
	return c.SendString("Hi World!")
}
