package account_types

import (
	accountTypeCorePort "github.com/codeclout/AccountEd/admin/internal/ports/core/account-types"
)

type UserAccountTypeAdminApiPort interface {
	CreateAccountType(in *string) (*accountTypeCorePort.NewAccountTypeOutput, error)
	RemoveAccountType(id *string) (*accountTypeCorePort.NewAccountTypeOutput, error)
	UpdateAccountType(accountType, accountTypeId *string) (*accountTypeCorePort.UpdatedAccountTypeOutput, error)
}
