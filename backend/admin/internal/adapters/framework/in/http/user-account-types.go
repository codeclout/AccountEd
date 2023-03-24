package http

import (
  "encoding/json"
  "errors"
  "fmt"
  "net/http"
  "strings"

  "github.com/codeclout/AccountEd/internal"
  "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
)

var v *validator.Validate

type UpdateAccountTypeInput struct {
  AccountTypeInput internal.AccountTypeInput
  CreateAccountTypeInput
}

type CreateAccountTypeInput struct {
  AccountType *string `json:"accountType" validate:"required,min=3"`
}

func (a *Adapter) initUserAccountTypeRoutes(app *fiber.App) *fiber.App {
  app.Delete("/user-account-type", a.processDeleteAccountType)
  app.Post("/user-account-type", a.processPostAccountType)
  app.Put("/user-account-type", a.processPutAccountType)

  return app
}

// processPostAccountType - validates account type POST requests as CreateAccountTypeInput
func (a *Adapter) processPostAccountType(ctx *fiber.Ctx) error {
  var t CreateAccountTypeInput

  payload := ctx.Body()
  if e := json.Unmarshal(payload, &t); e != nil {
    a.log("error", e.Error())

    _ = ctx.SendStatus(400)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         string(internal.ErrorInvalidJSON),
      ShouldRetry: internal.ShouldRetryRequest(400),
    })
  }

  v = validator.New()
  if e := v.Struct(t); e != nil {
    a.log("error", e.Error())

    _ = ctx.SendStatus(400)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         string(internal.ErrorFailedRequestValidation),
      ShouldRetry: internal.ShouldRetryRequest(400),
    })
  }

  result, e := a.HandleCreateAccountType(t.AccountType)
  if e != nil {
    if e.Error() == "duplicate" {
      _ = ctx.SendStatus(400)
    } else {
      _ = ctx.SendStatus(500)
    }

    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         fmt.Sprintf("%s | %s", e.Error(), *t.AccountType),
      ShouldRetry: internal.ShouldRetryRequest(ctx.Response().StatusCode()),
    })
  }

  return ctx.JSON(result)
}

func (a *Adapter) HandleCreateAccountType(in *string) (interface{}, error) {
  var f []byte

  result, e := a.userAccountTypeApi.CreateAccountType(in)
  if e != nil {
    a.log("error", e.Error())

    if strings.Contains(e.Error(), "duplicate key") {
      f = internal.ErrorDuplicateKey
    } else {
      f = internal.ErrorFailedAction
    }

    return nil, errors.New(string(f))
  }

  return result, nil
}

func (a *Adapter) processDeleteAccountType(ctx *fiber.Ctx) error {
  var t internal.AccountTypeInput

  id := ctx.Body()
  if e := json.Unmarshal(id, &t); e != nil {
    a.log("error", e.Error())

    _ = ctx.SendStatus(400)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         string(internal.ErrorInvalidJSON),
      ShouldRetry: internal.ShouldRetryRequest(400),
    })
  }

  v = validator.New()
  if e := v.Struct(t); e != nil {
    a.log("error", e.Error())

    _ = ctx.SendStatus(400)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         string(internal.ErrorFailedRequestValidation),
      ShouldRetry: internal.ShouldRetryRequest(400),
    })
  }

  result, e := a.HandleRemoveAccountType(t.Id)
  if e != nil {
    _ = ctx.SendStatus(500)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         e.Error(),
      ShouldRetry: internal.ShouldRetryRequest(500),
    })
  }

  return ctx.JSON(result)
}

func (a *Adapter) HandleRemoveAccountType(accountType *string) (interface{}, error) {
  result, e := a.userAccountTypeApi.RemoveAccountType(accountType)
  if e != nil {
    a.log("error", e.Error())
    return nil, errors.New(string(internal.ErrorFailedAction))
  }

  return result, nil
}

func (a *Adapter) processPutAccountType(ctx *fiber.Ctx) error {
  var t UpdateAccountTypeInput

  id := ctx.Body()
  e := json.Unmarshal(id, &t)
  if e != nil {
    a.log("error", e.Error())

    _ = ctx.SendStatus(http.StatusBadRequest)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         string(internal.ErrorInvalidJSON),
      ShouldRetry: internal.ShouldRetryRequest(http.StatusBadRequest),
    })
  }

  result, e := a.HandleUpdateAccountType(t.AccountType, t.AccountTypeInput.Id)
  if e != nil {
    _ = ctx.SendStatus(http.StatusInternalServerError)
    return ctx.JSON(internal.RequestErrorWithRetry{
      Msg:         e.Error(),
      ShouldRetry: internal.ShouldRetryRequest(http.StatusInternalServerError),
    })
  }

  return ctx.JSON(result)
}

func (a *Adapter) HandleUpdateAccountType(accountType, id *string) (interface{}, error) {
  result, e := a.userAccountTypeApi.UpdateAccountType(accountType, id)
  if e != nil {
    a.log("error", e.Error())
    return nil, errors.New(string(internal.ErrorFailedAction))
  }

  return result, nil
}
