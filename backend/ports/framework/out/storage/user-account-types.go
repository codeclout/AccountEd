package ports

type InsertID struct {
	InsertedID interface{}
}

type AccountTypeActionPort interface {
	GetAccountTypes(limit int64) ([]byte, error)
	InsertAccountType(acctType []byte) (InsertID, error)
	RemoveAccountType(id string) ([]byte, error)
	UpdateAccountType(in []byte) ([]byte, error)
	GetAccountTypeById(in []byte) ([]byte, error)
}
