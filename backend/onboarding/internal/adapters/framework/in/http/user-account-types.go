package http

import (
	"errors"
	"net/http"

	"github.com/codeclout/AccountEd/internal"
	"github.com/gofiber/fiber/v2"
)

// initUserAccountTypeRoutes - Registers handlers for the user account type routes
func (a *Adapter) initUserAccountTypeRoutes(app *fiber.App) *fiber.App {

	app.Get("/user-account-type", a.processGetAccountType)
	app.Get("/user-account-types", a.processListAccountTypes)

	return app
}

func (a *Adapter) processListAccountTypes(ctx *fiber.Ctx) error {
	q := ctx.Query("limit")

	limit := a.getRequestLimit(&q)
	result, e := a.HandleListAccountTypes(limit)
	if e != nil {
		_ = ctx.SendStatus(http.StatusInternalServerError)

		return ctx.JSON(internal.RequestErrorWithRetry{
			Msg:         e.Error(),
			ShouldRetry: internal.ShouldRetryRequest(500),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleListAccountTypes(in *int16) (interface{}, error) {
	result, e := a.userAccountTypeApi.GetAccountTypes(in)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(internal.ErrorFailedAction))
	}

	return result, nil
}

func (a *Adapter) processGetAccountType(ctx *fiber.Ctx) error {
	var t internal.AccountTypeInput

	if e := ctx.QueryParser(&t); e != nil {
		a.log("error", e.Error())

		_ = ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(internal.RequestErrorWithRetry{
			Msg:         string(internal.ErrorInvalidJSON),
			ShouldRetry: internal.ShouldRetryRequest(http.StatusBadRequest),
		})
	}

	result, e := a.HandleFetchAccountType(t.Id)
	if e != nil {
		_ = ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(internal.RequestErrorWithRetry{
			Msg:         e.Error(),
			ShouldRetry: internal.ShouldRetryRequest(http.StatusBadRequest),
		})
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleFetchAccountType(id *string) (interface{}, error) {
	result, e := a.userAccountTypeApi.FetchAccountType(id)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(internal.ErrorFailedAction))
	}

	return result, nil
}
