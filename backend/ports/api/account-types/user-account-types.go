package account_types

import (
	ports "github.com/codeclout/AccountEd/ports/core/account-types"
)

type UserAccountTypeApiPort interface {
	CreateAccountType(in string) (ports.NewAccountTypeOutput, error)
	GetAccountTypes(limit int64) ([]ports.NewAccountTypeOutput, error)
	RemoveAccountType(id string) (ports.NewAccountTypeOutput, error)
	UpdateAccountType(in []byte) (ports.NewAccountTypeOutput, error)
}
