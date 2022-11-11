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

	accountType.Post("/user-account-type", a.HandlePostAccountType)
	accountType.Get("/user-account-types", a.HandleGetAccountTypes)

	return accountType
}

func (a *Adapter) HandlePostAccountType(c *fiber.Ctx) error {
	return a.HandleCreateAccountType(c)
}

func (a *Adapter) HandleCreateAccountType(i interface{}) error {
	c := i.(*fiber.Ctx)

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

func (a *Adapter) HandleGetAccountTypes(c *fiber.Ctx) error {
	return a.HandleListAccountTypes(c)
}

func (a *Adapter) HandleListAccountTypes(i interface{}) error {
	var lmt int64

	c := i.(*fiber.Ctx)
	q := c.Query("limit")

	lmt = a.getRequestLimit(q)

	if result, e := a.api.GetAccountTypes(lmt); e != nil {
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
