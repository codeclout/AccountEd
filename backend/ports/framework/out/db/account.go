package ports

type AccountDbPort interface {
	CloseConnection()
	InsertAccountType(handle string, acctType string) (InsertID, error)
}

type InsertID struct {
	InsertedID interface{}
}
