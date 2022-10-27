package http

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateAccountTypeInput struct {
	AccountType string `json:"accountType" validate:"required,min=3"`
}

var validate *validator.Validate

func (a *Adapter) initUserRoutes() *fiber.App {
	var accountType = fiber.New()

	accountType.Post("/user-account-type", a.HandleCreateAccountType)

	return accountType
}

func (a *Adapter) HandleCreateAccountType(c *fiber.Ctx) error {
	var payload = c.Body()
	var t CreateAccountTypeInput

	if e := json.Unmarshal(payload, &t); e != nil {
		a.log("error", e.Error())
		return c.SendStatus(400)
	}

	cat := CreateAccountTypeInput{AccountType: t.AccountType}
	validate = validator.New()

	if e := validate.Struct(cat); e != nil {
		a.log("error", e.Error())
		return c.SendStatus(400)
	}

	result, e := a.api.CreateAccountType(t.AccountType)

	if e != nil {
		a.log("error", e.Error())
		return c.SendStatus(400)
	}

	return c.JSON(result)
}

func (a *Adapter) GetAccountTypes() error {
	return nil
}
