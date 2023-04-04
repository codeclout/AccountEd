package storage

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type AccountTypeActionPort interface {
	GetAccountTypes(ctx context.Context, limit int16) (*[]internal.AccountTypeOut, error)
	GetAccountTypeById(ctx context.Context, in internal.AccountTypeIn) (*internal.AccountTypeOut, error)
}
