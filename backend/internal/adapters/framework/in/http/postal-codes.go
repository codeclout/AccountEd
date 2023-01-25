package http

import (
  "net/http"

  "github.com/codeclout/AccountEd/internal/adapters/framework/in/http/requests"
  "github.com/gofiber/fiber/v2"
)

func (a *Adapter) initPostalCodeRoutes(app *fiber.App) *fiber.App {
  app.Get("/postal-code", a.processGetPostalCodeDetails)

  return app
}

func (a *Adapter) processGetPostalCodeDetails(ctx *fiber.Ctx) error {
  q := ctx.Query("address")

  result, e := a.HandleFetchPostalCodeDetails(&q)
  if e != nil {
    _ = ctx.SendStatus(http.StatusInternalServerError)

    return ctx.JSON(requests.RequestErrorWithRetry{
      Msg:         e.Error(),
      ShouldRetry: requests.ShouldRetryRequest(500),
    })
  }

  return ctx.JSON(result)
}

func (a *Adapter) HandleFetchPostalCodeDetails(address *string) (interface{}, error) {
  return a.postalCodeApi.FetchPostalCodeDetails(address)
}
