package http

import (
  "net/http"

  "github.com/codeclout/AccountEd/onboarding/internal/adapters/framework/in/http/requests"
  "github.com/gofiber/fiber/v2"
)

func (a *in.Adapter) initPostalCodeRoutes(app *fiber.App) *fiber.App {
  app.Get("/postal-code", a.processGetPostalCodeDetails)

  return app
}

func (a *in.Adapter) processGetPostalCodeDetails(ctx *fiber.Ctx) error {
  q := ctx.Query("address")

  result, e := a.HandleFetchPostalCodeDetails(&q)
  if e != nil {
    if e.Error() == "ZERO_RESULTS" || e.Error() == "INVALID_REQUEST" {
      _ = ctx.SendStatus(http.StatusNotFound)

      return ctx.JSON(requests.RequestErrorWithRetry{
        Msg:         e.Error(),
        ShouldRetry: requests.ShouldRetryRequest(http.StatusNotFound),
      })

    }

    _ = ctx.SendStatus(http.StatusInternalServerError)

    return ctx.JSON(requests.RequestErrorWithRetry{
      Msg:         e.Error(),
      ShouldRetry: requests.ShouldRetryRequest(http.StatusInternalServerError),
    })
  }

  return ctx.JSON(result)
}

func (a *in.Adapter) HandleFetchPostalCodeDetails(address *string) (interface{}, error) {
  return a.postalCodeApi.FetchPostalCodeDetails(address)
}
