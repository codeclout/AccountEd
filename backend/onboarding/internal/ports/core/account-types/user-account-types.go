package account_types

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type UserAccountTypeCorePort interface {
	ListAccountTypes(ctx context.Context, limit int16) (*[]internal.AccountTypeOut, error)
	FetchAccountType(ctx context.Context, in internal.AccountTypeIn) (*internal.AccountTypeOut, error)
}
