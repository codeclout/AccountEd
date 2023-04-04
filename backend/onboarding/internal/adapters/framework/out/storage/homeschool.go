package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/codeclout/AccountEd/pkg/storage"
)

var registrationCollectionName = []byte("account")

func (a *Adapter) Login() {}

func (a *Adapter) Register(ctx context.Context, data *[]interface{}) ([]storage.MongoInsertOutput, error) {
	var out []storage.MongoInsertOutput

	collection := a.mongoActions.Db.Collection(string(registrationCollectionName))
	t := a.mongoActions.GetTimeStamp()

	account, e := collection.InsertMany(ctx, *data)
	if e != nil {
		a.log("error", e.Error())

		if mongoError, ok := e.(mongo.BulkWriteException); ok {
			a.mongoActions.BulkWriteError.BulkWriteErrorExceptionMsg.WriteErrors = mongoError.WriteErrors
			return nil, errors.New(a.mongoActions.BulkWriteError.Error())
		}

		return nil, e
	}

	for _, o := range account.InsertedIDs {
		x := o.(primitive.ObjectID).Hex()

		y := storage.MongoInsertOutput{
			Id:        &x,
			TimeStamp: &t,
		}

		_ = append(out, y)
	}

	return out, nil

}
