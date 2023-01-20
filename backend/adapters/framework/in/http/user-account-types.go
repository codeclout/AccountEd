package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/codeclout/AccountEd/adapters/framework/in/http/requests"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var v *validator.Validate

type AccountTypeInput struct {
	Id *string `json:"id" validate:"required"`
}

type CreateAccountTypeInput struct {
	AccountType *string `json:"accountType" validate:"required,min=3"`
}

type UpdateAccountTypeInput struct {
	AccountTypeInput
	CreateAccountTypeInput
}

// initUserRoutes - Registers handlers for the user account type routes
func (a *Adapter) initUserRoutes() *fiber.App {
	var accountType = fiber.New()

	accountType.Delete("/user-account-type", a.processDeleteAccountType)
	accountType.Get("/user-account-type", a.processGetAccountType)
	accountType.Get("/user-account-types", a.processListAccountTypes)
	accountType.Post("/user-account-type", a.processPostAccountType)
	accountType.Put("/user-account-type", a.processPutAccountType)

	return accountType
}

func (a *Adapter) processPostAccountType(ctx *fiber.Ctx) error {
	var t CreateAccountTypeInput

	payload := ctx.Body()
	if e := json.Unmarshal(payload, &t); e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	v = validator.New()
	if e := v.Struct(t); e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedRequestValidation),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	result, e := a.HandleCreateAccountType(t.AccountType)
	if e != nil {
		if e.Error() == "duplicate" {
			_ = ctx.SendStatus(400)
		} else {
			_ = ctx.SendStatus(500)
		}

		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         fmt.Sprintf("%s | %s", e.Error(), *t.AccountType),
			ShouldRetry: requests.ShouldRetryRequest(ctx.Response().StatusCode()),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleCreateAccountType(in *string) (interface{}, error) {
	var f []byte

	result, e := a.api.CreateAccountType(in)
	if e != nil {
		a.log("error", e.Error())

		if strings.Contains(e.Error(), "duplicate key") {
			f = requests.ErrorDuplicateKey
		} else {
			f = requests.ErrorFailedAction
		}

		return nil, errors.New(string(f))
	}

	return result, nil
}

func (a *Adapter) processListAccountTypes(ctx *fiber.Ctx) error {
	q := ctx.Query("limit")

	limit := a.getRequestLimit(&q)
	result, e := a.HandleListAccountTypes(limit)
	if e != nil {
		_ = ctx.SendStatus(500)

		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         e.Error(),
			ShouldRetry: requests.ShouldRetryRequest(500),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleListAccountTypes(in *int16) (interface{}, error) {
	result, e := a.api.GetAccountTypes(in)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(requests.ErrorFailedAction))
	}

	return result, nil
}

func (a *Adapter) processDeleteAccountType(ctx *fiber.Ctx) error {
	var t AccountTypeInput

	id := ctx.Body()
	if e := json.Unmarshal(id, &t); e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	v = validator.New()
	if e := v.Struct(t); e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorFailedRequestValidation),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	result, e := a.HandleRemoveAccountType(t.Id)
	if e != nil {
		_ = ctx.SendStatus(500)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         e.Error(),
			ShouldRetry: requests.ShouldRetryRequest(500),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleRemoveAccountType(accountType *string) (interface{}, error) {
	result, e := a.api.RemoveAccountType(accountType)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(requests.ErrorFailedAction))
	}

	return result, nil
}

func (a *Adapter) processPutAccountType(ctx *fiber.Ctx) error {
	var t UpdateAccountTypeInput

	id := ctx.Body()
	e := json.Unmarshal(id, &t)
	if e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	result, e := a.HandleUpdateAccountType(t.AccountType, t.Id)
	if e != nil {
		_ = ctx.SendStatus(500)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         e.Error(),
			ShouldRetry: requests.ShouldRetryRequest(500),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleUpdateAccountType(accountType, id *string) (interface{}, error) {
	result, e := a.api.UpdateAccountType(accountType, id)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(requests.ErrorFailedAction))
	}

	return result, nil
}

func (a *Adapter) processGetAccountType(ctx *fiber.Ctx) error {
	var t AccountTypeInput

	if e := ctx.QueryParser(&t); e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         string(requests.ErrorInvalidJSON),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	result, e := a.HandleFetchAccountType(t.Id)
	if e != nil {
		_ = ctx.SendStatus(400)
		return ctx.JSON(requests.RequestErrorWithRetry{
			Msg:         e.Error(),
			ShouldRetry: requests.ShouldRetryRequest(400),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleFetchAccountType(id *string) (interface{}, error) {
	result, e := a.api.FetchAccountType(id)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(requests.ErrorFailedAction))
	}

	return result, nil
}
