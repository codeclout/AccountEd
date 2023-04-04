package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInsertOutput struct {
	Id        *string
	TimeStamp *primitive.DateTime
}
