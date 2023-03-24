package storage

type AccountTypeActionPort interface {
	GetAccountTypes(limit *int16) (*[]byte, error)
	GetAccountTypeById(id *string) (*[]byte, error)
}
