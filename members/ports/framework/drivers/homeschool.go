package drivers

import (
  "context"

  "github.com/gofiber/fiber/v2"

  mt "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolDriverPort interface {
  HandlePreRegistration(ctx context.Context, in *mt.PrimaryMemberStartRegisterIn) (*mt.PrimaryMemberStartRegisterOut, error)
  InitializeAPI(app *fiber.App) []*fiber.App
}
