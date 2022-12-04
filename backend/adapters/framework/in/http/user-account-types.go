package http

import (
	"encoding/json"

	"github.com/codeclout/AccountEd/adapters/framework/in/http/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var v *validator.Validate

type AccountTypeInput struct {
	Id string `json:"id" validate:"required"`
}

type CreateAccountTypeInput struct {
	AccountType string `json:"accountType" validate:"required,min=3"`
}

type UpdateAccountTypeInput struct {
	AccountTypeInput
	CreateAccountTypeInput
}

// initUserRoutes - Registers methods for the user account type routes
func (a *Adapter) initUserRoutes() *fiber.App {
	var accountType = fiber.New()

	accountType.Post("/user-account-type", a.handlePostAccountType)
	accountType.Get("/user-account-types", a.HandleGetAccountTypes)
	accountType.Delete("/user-account-type", a.HandleDeleteAccountType)
	accountType.Put("/user-account-type", a.HandlePutAccountType)

	return accountType
}

// HandlePostAccountType - wraps the handler to create a new account type
func (a *Adapter) handlePostAccountType(c *fiber.Ctx) error {
	return a.handleCreateAccountType(c)
}

func (a *Adapter) handleCreateAccountType(i interface{}) error {
	c := i.(*fiber.Ctx)

	var payload = c.Body()
	var t CreateAccountTypeInput

	if e := json.Unmarshal(payload, &t); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	cat := CreateAccountTypeInput{AccountType: t.AccountType}
	v = validator.New()

	if e := v.Struct(cat); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedRequestValidation),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	result, e := a.api.CreateAccountType(t.AccountType)

	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedAction),
			ShouldRetry: requests.ShouldRetryRequest(500),
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
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedAction),
			ShouldRetry: requests.ShouldRetryRequest(500),
		})
	} else {
		return c.JSON(result)
	}
}

func (a *Adapter) HandleDeleteAccountType(c *fiber.Ctx) error {
	return a.HandleRemoveAccountType(c)
}

func (a *Adapter) HandleRemoveAccountType(i interface{}) error {
	var t AccountTypeInput

	c := i.(*fiber.Ctx)
	id := c.Body()

	if e := json.Unmarshal(id, &t); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	cat := AccountTypeInput{Id: t.Id}
	v = validator.New()

	if e := v.Struct(cat); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedRequestValidation),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	result, e := a.api.RemoveAccountType(t.Id)

	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedAction),
			ShouldRetry: requests.ShouldRetryRequest(500),
		})
	}

	return c.JSON(result)
}

func (a *Adapter) HandlePutAccountType(c *fiber.Ctx) error {
	id := c.Body()
	return a.HandleUpdateAccountType(id, c)
}

func (a *Adapter) HandleUpdateAccountType(id []byte, i interface{}) error {
	var t UpdateAccountTypeInput

	c := i.(*fiber.Ctx)
	e := json.Unmarshal(id, &t)

	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	b, e := json.Marshal(t)
	r, e := a.api.UpdateAccountType(b)

	return c.JSON(r)
}
