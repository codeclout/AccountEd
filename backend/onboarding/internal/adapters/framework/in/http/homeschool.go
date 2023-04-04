package http

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	app.Post("/register", timeout.New(a.processRegistration, 2*time.Second))
	app.Post("/login", a.processLogin)

	return app
}

func (a *Adapter) processRegistration(ctx *fiber.Ctx) error {
	var registerHomeSchool internal.HomeSchoolRegisterIn
	payload := ctx.Body()

	if e := json.Unmarshal(payload, &registerHomeSchool); e != nil {
		a.log("error", e.Error())
		return ctx.JSON(ctx.SendStatus(400))
	}

	for _, v := range registerHomeSchool.ParentGuardians {
		internal.ValidateHomeschoolRegistration(v)
	}

	result, e := a.HandleRegistration(ctx.UserContext(), registerHomeSchool)
	if e != nil {
		return ctx.JSON(e)
	}

	return ctx.JSON(result)
}

func (a *Adapter) HandleRegistration(ctx context.Context, in internal.HomeSchoolRegisterIn) (internal.HomeSchoolRegisterOut, error) {

	ch := make(chan internal.HomeSchoolRegisterOut, 1)
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	e := a.homeschoolOnboardApi.RegisterHomeSchool(ctx, ch, in)
	if e != nil {
		return internal.HomeSchoolRegisterOut{}, e
	}

	select {

	case <-ctx.Done():
		a.log("info", "")
		return internal.HomeSchoolRegisterOut{}, ctx.Err()

	case out := <-ch:
		a.log("info", "")
		return out, ctx.Err()
	}

}

func (a *Adapter) processLogin(ctx *fiber.Ctx) error {
	return nil
}

func (a *Adapter) HandleLogin(ctx context.Context) error {
	return nil
}
