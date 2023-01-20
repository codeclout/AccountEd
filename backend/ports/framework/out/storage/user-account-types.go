package ports

import "go.mongodb.org/mongo-driver/bson/primitive"

type InsertId string

type InsertAccountOutput struct {
	InsertId  *InsertId
	TimeStamp *primitive.DateTime
}

type AccountTypeActionPort interface {
	GetAccountTypes(limit *int16) (*[]byte, error)
	InsertAccountType(acctType *[]byte) (*[]byte, error)
	RemoveAccountType(id *string) (*[]byte, error)
	UpdateAccountType(accountTypeName, accountTypeId *string) (*int64, error)
	GetAccountTypeById(id *string) (*[]byte, error)
}
