package drivers

import (
  "context"

  "github.com/gofiber/fiber/v2"

  membertypes "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolDriverPort interface {
  //HandleParentGuardiansByAccountId(ctx context.Context, id string) *membertypes.ParentGuardianOut
  HandleRegistration(ctx context.Context, in *membertypes.HomeSchoolRegisterIn) (*membertypes.HomeSchoolRegisterOut, error)
  InitializeAPI(app *fiber.App) []*fiber.App
}
