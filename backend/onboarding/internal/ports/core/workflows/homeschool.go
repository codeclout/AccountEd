package workflows

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type HomeSchoolCorePort interface {
	LogIn(ctx context.Context, in LogIn)
	Register(ctx context.Context, in *internal.HomeSchoolRegisterIn) (*internal.HomeSchoolRegisterOut, error)
	UsersByAccountId(ctx context.Context)
	UserById(ctx context.Context)
	UserByUsername(ctx context.Context)
	// TODO - add exception handler
}

type LogIn struct {
	Pin      string `json:"pin"`
	Username string `json:"username"`
}
