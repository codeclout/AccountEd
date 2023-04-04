package http

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type UserAccountPort interface {
	HandleListAccountTypes(ctx context.Context, in int16) (*[]internal.AccountTypeOut, error)
	HandleFetchAccountType(ctx context.Context, in internal.AccountTypeIn) (*internal.AccountTypeOut, error)
}
