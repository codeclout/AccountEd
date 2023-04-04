package workflows

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type OnboardHomeschoolApiPort interface {
	LoginHomeSchool(ctx context.Context)
	RegisterHomeSchool(ctx context.Context, ch chan internal.HomeSchoolRegisterOut, in internal.HomeSchoolRegisterIn) error
}
