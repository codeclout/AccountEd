package core

import (
	"context"

	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

type Authentication interface {
	AuthenticationOptions(ctx context.Context, id string) (*sessiontypes.AuthenticationOptions, error)
	RegistrationOptions(ctx context.Context, id string) (*sessiontypes.RegistrationOptions, error)
}
