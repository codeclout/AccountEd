package ports

type UserAccountTypeDbPort interface {
	CloseConnection()
	InsertAccountType(acctType []byte) (InsertID, error)
	GetAccountTypes(limit int64) ([]byte, error)
}

type InsertID struct {
	InsertedID interface{}
}
