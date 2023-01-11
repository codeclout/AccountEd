package ports

import "go.mongodb.org/mongo-driver/bson/primitive"

type InsertId string

type InsertAccountOutput struct {
	InsertId  *InsertId
	TimeStamp *primitive.DateTime
}

type AccountTypeActionPort interface {
	GetAccountTypes(limit int64) ([]byte, error)
	InsertAccountType(acctType []byte) (*[]byte, error)
	RemoveAccountType(id string) ([]byte, error)
	UpdateAccountType(in []byte) ([]byte, error)
	GetAccountTypeById(in []byte) ([]byte, error)
}
