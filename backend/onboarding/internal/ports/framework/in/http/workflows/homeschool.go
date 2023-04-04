package workflows

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type UserAccountPort interface {
	HandleRegistration(ctx context.Context, in internal.HomeSchoolRegisterIn) (string, error)
	HandleLogin(ctx context.Context) error
}
