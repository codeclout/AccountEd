package ports

type UserAccountTypeDbPort interface {
	CloseConnection()
	GetAccountTypes(limit int64) ([]byte, error)
	InsertAccountType(acctType []byte) (InsertID, error)
	RemoveAccountType(id string) ([]byte, error)
	UpdateAccountType(in []byte) ([]byte, error)
	GetAccountTypeById(in []byte) ([]byte, error)
}

type InsertID struct {
	InsertedID interface{}
}
