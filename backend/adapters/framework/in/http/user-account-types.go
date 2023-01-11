package http

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/codeclout/AccountEd/adapters/framework/in/http/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var v *validator.Validate

type AccountTypeInput struct {
	Id string `json:"id" validate:"required"`
}

type CreateAccountTypeInput struct {
	AccountType *string `json:"accountType" validate:"required,min=3"`
}

type UpdateAccountTypeInput struct {
	AccountTypeInput
	CreateAccountTypeInput
}

// initUserRoutes - Registers methods for the user account type routes
func (a *Adapter) initUserRoutes() *fiber.App {
	var accountType = fiber.New()

	accountType.Delete("/user-account-type", a.handleDeleteAccountType)
	accountType.Get("/user-account-type", a.handleGetAccountType)
	accountType.Get("/user-account-types", a.handleGetAccountTypes)
	accountType.Post("/user-account-type", a.handlePostAccountType)
	accountType.Put("/user-account-type", a.handlePutAccountType)

	return accountType
}

// HandlePostAccountType - wraps the handler to create a new account type
func (a *Adapter) handlePostAccountType(c *fiber.Ctx) error {
	return a.HandleCreateAccountType(c)
}

func (a *Adapter) HandleCreateAccountType(i interface{}) error {
	var f string
	var t CreateAccountTypeInput

	c := i.(*fiber.Ctx)
	var payload = c.Body()

	if e := json.Unmarshal(payload, &t); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	v = validator.New()

	if e := v.Struct(t); e != nil {
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

		if strings.Contains(e.Error(), "duplicate key") {
			_ = c.SendStatus(400)
			f = string(requests.ErrorDuplicateKey)
		} else {
			_ = c.SendStatus(500)
			f = string(requests.ErrorFailedAction)
		}

		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         fmt.Sprintf("%s | %s", f, *t.AccountType),
			ShouldRetry: requests.ShouldRetryRequest(c.Response().StatusCode()),
		})
	}

	return c.JSON(result)
}

func (a *Adapter) handleGetAccountTypes(c *fiber.Ctx) error {
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

func (a *Adapter) handleDeleteAccountType(c *fiber.Ctx) error {
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

	v = validator.New()

	if e := v.Struct(t); e != nil {
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

func (a *Adapter) handlePutAccountType(c *fiber.Ctx) error {
	id := c.Body()
	return a.HandleUpdateAccountType(id, c)
}

func (a *Adapter) HandleUpdateAccountType(id []byte, i interface{}) error {
	var t UpdateAccountTypeInput

	c := i.(*fiber.Ctx)
	json.Valid(id)

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
	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	r, e := a.api.UpdateAccountType(b)
	if e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(500)
		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedAction),
			ShouldRetry: requests.ShouldRetryRequest(500),
		})
	}

	return c.JSON(r)
}

func (a *Adapter) handleGetAccountType(c *fiber.Ctx) error {
	return a.HandleFetchAccountType(c)
}

func (a *Adapter) HandleFetchAccountType(i interface{}) error {
	var t AccountTypeInput

	c := i.(*fiber.Ctx)

	if e := c.QueryParser(&t); e != nil {
		a.log("error", e.Error())

		_ = c.SendStatus(400)

		return c.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	b, _ := json.Marshal(t)
	r, e := a.api.FetchAccountType(b)

	if e != nil {
		a.log("error", e.Error())
		return e
	}

	return c.JSON(r)
}
