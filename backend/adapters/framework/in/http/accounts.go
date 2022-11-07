package http

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

type CreateAccountTypeInput struct {
	AccountType string `json:"accountType" validate:"required,min=3"`
}

func (a *Adapter) initUserRoutes() *fiber.App {
	var accountType = fiber.New()

	accountType.Post("/user-account-type", a.HandleCreateAccountType)
	accountType.Get("/user-account-types", a.HandleListAccountTypes)

	return accountType
}

func (a *Adapter) HandleCreateAccountType(c *fiber.Ctx) error {
	var payload = c.Body()
	var t CreateAccountTypeInput

	if e := json.Unmarshal(payload, &t); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(RequestErrorWithRetry{
			Msg:         string(ErrorInvalidJSON),
			ShouldRetry: ShouldRetryRequest(400),
		})
	}

	cat := CreateAccountTypeInput{AccountType: t.AccountType}
	validate = validator.New()

	if e := validate.Struct(cat); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(RequestErrorWithRetry{
			Msg:         string(ErrorFailedRequestValidation),
			ShouldRetry: ShouldRetryRequest(400),
		})
	}

	result, e := a.api.CreateAccountType(t.AccountType)

	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(RequestErrorWithRetry{
			Msg:         string(ErrorFailedAction),
			ShouldRetry: ShouldRetryRequest(500),
		})
	}

	return c.JSON(result)
}

func (a *Adapter) HandleListAccountTypes(c *fiber.Ctx) error {
	if result, e := a.api.GetAccountTypes("account_type"); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(RequestErrorWithRetry{
			Msg:         string(ErrorFailedAction),
			ShouldRetry: ShouldRetryRequest(500),
		})
	} else {
		return c.JSON(result)
	}
}
