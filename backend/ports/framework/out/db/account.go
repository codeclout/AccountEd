package ports

type AccountDbPort interface {
	CloseConnection()
	InsertAccountType(collection string, acctType []byte) (InsertID, error)
	GetAccountTypes(collection string) ([]byte, error)
}

type InsertID struct {
	InsertedID interface{}
}
