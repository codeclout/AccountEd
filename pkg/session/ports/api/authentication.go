package api

import (
  "context"

  sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

type Authenticaiton interface {
  GetAuthenticationOptions(ctx context.Context, ch chan<- sessiontypes.AuthenticationOptions, id string) error
  GetRegistrationOptions(ctx context.Context, ch chan<- sessiontypes.RegistrationOptions, id string) error
}
