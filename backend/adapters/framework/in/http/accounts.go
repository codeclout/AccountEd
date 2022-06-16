package http

import "github.com/gofiber/fiber/v2"

type CreateAccountTypeInput struct {
	AccountType string `json:"accountType"`
}

func (a Adapter) CreateAccountType(c *fiber.Ctx) error {
	cat := new(CreateAccountTypeInput)

	if e := c.BodyParser(cat); e != nil {
		return e
	}

	result, e := a.api.CreateAccountType(cat.AccountType)

	if e != nil {
		return e
	}

	return c.JSON(result)
}
