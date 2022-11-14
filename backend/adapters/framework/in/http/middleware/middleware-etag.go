package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AirCollision412() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var r = c.Response().Header.Header()
		h := c.GetReqHeaders()

		for header, value := range h {
			fmt.Println(header, value)
		}

		s := c.Get("If-Match")
		fmt.Println(s)
		fmt.Println(r)

		return c.Next()
	}
}
