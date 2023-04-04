package http

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/middleware/timeout"

	"github.com/codeclout/AccountEd/onboarding/internal"
	"github.com/codeclout/AccountEd/pkg"

	"github.com/gofiber/fiber/v2"
)

// initUserAccountTypeRoutes - Registers handlers for the user account type routes
func (a *Adapter) initUserAccountTypeRoutes(app *fiber.App) *fiber.App {
	app.Get("/user-account-type", timeout.New(a.processGetAccountType, 1*time.Second))
	app.Get("/user-account-types", timeout.New(a.processListAccountTypes, 1*time.Second))

	return app
}

func handleErrorWithRetry(ctx *fiber.Ctx, e error, status int) error {
	return ctx.JSON(pkg.RequestErrorWithRetry{
		Msg:         e.Error(),
		ShouldRetry: pkg.ShouldRetryRequest(status),
	})
}

func (a *Adapter) processListAccountTypes(ctx *fiber.Ctx) error {
	q := ctx.Query("limit")
	limit := a.getRequestLimit(&q)

	result, e := a.HandleListAccountTypes(ctx.UserContext(), *limit)
	if e != nil {
		a.log("error", e.Error())
		ctx.Status(fiber.StatusInternalServerError)

		return handleErrorWithRetry(ctx, e, fiber.StatusInternalServerError)
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleListAccountTypes(ctx context.Context, in int16) (*[]internal.AccountTypeOut, error) {

	ch := make(chan *[]internal.AccountTypeOut, 1)
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	e := a.userAccountTypeApi.GetAccountTypes(ctx, in, ch)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(pkg.ErrorFailedAction))
	}

	select {
	case <-ctx.Done():
		a.log("info", "<Transaction ID> List AccountTypes Timeout")
		return nil, ctx.Err()

	case out := <-ch:
		a.log("info", "Transaction ID not implemented - List AccountTypes")
		return out, ctx.Err()
	}
}

func (a *Adapter) processGetAccountType(ctx *fiber.Ctx) error {
	var t internal.AccountTypeIn

	if e := ctx.QueryParser(&t); e != nil {
		a.log("error", e.Error())
		ctx.Status(fiber.StatusBadRequest)
		return handleErrorWithRetry(ctx, e, fiber.StatusBadRequest)
	}

	if ok, _ := internal.ValidateAccountTypeId(t); !ok {
		ctx.Status(fiber.StatusForbidden)
		return handleErrorWithRetry(ctx, errors.New("forbidden"), fiber.StatusForbidden)
	}

	result, e := a.HandleFetchAccountType(ctx.UserContext(), t)
	if e != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return handleErrorWithRetry(ctx, e, fiber.StatusInternalServerError)
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleFetchAccountType(ctx context.Context, in internal.AccountTypeIn) (*internal.AccountTypeOut, error) {

	ch := make(chan *internal.AccountTypeOut, 1)
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	e := a.userAccountTypeApi.FetchAccountType(ctx, in, ch)
	if e != nil {
		a.log("error", e.Error())
		return nil, errors.New(string(pkg.ErrorFailedAction))
	}

	select {
	case <-ctx.Done():
		a.log("error", "<Transaction ID> By ID AccountType Timeout")
		return nil, ctx.Err()

	case out := <-ch:
		a.log("info", "Transaction ID not implemented - By ID AccountType")
		return out, ctx.Err()
	}
}
