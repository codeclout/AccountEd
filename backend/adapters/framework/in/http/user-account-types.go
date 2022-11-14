package http

import (
	"encoding/json"

	"github.com/codeclout/AccountEd/adapters/framework/in/http/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var v *validator.Validate

type CreateAccountTypeInput struct {
	AccountType string `json:"accountType" validate:"required,min=3"`
}

type AccountTypeInput struct {
	Id string `json:"id" validate:"required"`
}

func (a *Adapter) initUserRoutes() *fiber.App {
	var accountType = fiber.New()

	accountType.Post("/user-account-type", a.HandlePostAccountType)
	accountType.Get("/user-account-types", a.HandleGetAccountTypes)
	accountType.Delete("/user-account-type", a.HandleDeleteAccountType)

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
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorInvalidJSON),
			ShouldRetry: helpers.ShouldRetryRequest(400),
		})
	}

	cat := CreateAccountTypeInput{AccountType: t.AccountType}
	v = validator.New()

	if e := v.Struct(cat); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorFailedRequestValidation),
			ShouldRetry: helpers.ShouldRetryRequest(400),
		})
	}

	result, e := a.api.CreateAccountType(t.AccountType)

	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorFailedAction),
			ShouldRetry: helpers.ShouldRetryRequest(500),
		})
	}

	return c.JSON(result)
}

func (a *Adapter) HandleGetAccountTypes(c *fiber.Ctx) error {
	q := c.Query("limit")
	limit := a.getRequestLimit(q)

	return a.HandleListAccountTypes(limit, c)
}

func (a *Adapter) HandleListAccountTypes(limit int64, i interface{}) error {
	c := i.(*fiber.Ctx)

	if result, e := a.api.GetAccountTypes(limit); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorFailedAction),
			ShouldRetry: helpers.ShouldRetryRequest(500),
		})
	} else {
		return c.JSON(result)
	}
}

func (a *Adapter) HandleDeleteAccountType(c *fiber.Ctx) error {
	return a.HandleRemoveAccountType(c)
}

func (a *Adapter) HandleRemoveAccountType(i interface{}) error {
	c := i.(*fiber.Ctx)
	id := c.Body()

	var t AccountTypeInput

	if e := json.Unmarshal(id, &t); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorInvalidJSON),
			ShouldRetry: helpers.ShouldRetryRequest(400),
		})
	}

	cat := AccountTypeInput{Id: t.Id}
	v = validator.New()

	if e := v.Struct(cat); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorFailedRequestValidation),
			ShouldRetry: helpers.ShouldRetryRequest(400),
		})
	}

	result, e := a.api.RemoveAccountType(t.Id)

	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(helpers.RequestErrorWithRetry{
			Msg:         string(helpers.ErrorFailedAction),
			ShouldRetry: helpers.ShouldRetryRequest(500),
		})
	}

	return c.JSON(result)
}
