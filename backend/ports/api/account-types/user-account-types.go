package account_types

import (
	ports "github.com/codeclout/AccountEd/ports/core/account-types"
)

type UserAccountTypeApiPort interface {
	CreateAccountType(in *string) (*ports.NewAccountTypeOutput, error)
	GetAccountTypes(limit *int16) (*[]ports.NewAccountTypeOutput, error)
	RemoveAccountType(id *string) (*ports.NewAccountTypeOutput, error)
	UpdateAccountType(accountType, accountTypeId *string) (*ports.UpdatedAccountTypeOutput, error)
	FetchAccountType(id *string) (*ports.NewAccountTypeOutput, error)
}
