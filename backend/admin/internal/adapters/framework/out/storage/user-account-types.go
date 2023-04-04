package storage

type AccountTypeActionPort interface {
	InsertAccountType(acctType *[]byte) (*[]byte, error)
	RemoveAccountType(id *string) (*[]byte, error)
	UpdateAccountType(accountTypeName, accountTypeId *string) (*int64, error)
}
