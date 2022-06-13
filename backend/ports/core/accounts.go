package ports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountPort interface {
	NewAccountType(in string) (AccountTypeOutput, error)
}

type AccountTypeOutput struct {
	_id         primitive.ObjectID
	accountType string
}
