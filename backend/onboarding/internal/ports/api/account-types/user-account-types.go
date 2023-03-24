package account_types

import (
	ports "github.com/codeclout/AccountEd/onboarding/internal/ports/core/account-types"
)

type UserAccountTypeApiPort interface {
	GetAccountTypes(limit *int16) (*[]ports.NewAccountTypeOutput, error)
	FetchAccountType(id *string) (*ports.NewAccountTypeOutput, error)
}
