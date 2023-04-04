package account_types

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

type UserAccountTypeApiPort interface {
	GetAccountTypes(ctx context.Context, limit int16, ch chan *[]internal.AccountTypeOut) error
	FetchAccountType(ctx context.Context, in internal.AccountTypeIn, ch chan *internal.AccountTypeOut) error
}
