package ports

type AccountDbPort interface {
	CloseConnection()
	InsertAccountType(collection string, acctType []byte) (InsertID, error)
}

type InsertID struct {
	InsertedID interface{}
}
