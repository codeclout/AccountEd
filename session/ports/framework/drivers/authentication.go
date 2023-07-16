package driver

import (
	"context"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Authentication interface {
	HandleAuthenticationOptions(ctx context.Context, id string) sessiontypes.AuthenticationOptions
	HandleRegistrationOptions(ctx context.Context, id string) sessiontypes.RegistrationOptions
}
