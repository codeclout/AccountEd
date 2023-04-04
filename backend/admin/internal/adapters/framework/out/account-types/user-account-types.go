package account_types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/codeclout/AccountEd/admin/internal/adapters/framework/out/storage"
	"github.com/codeclout/AccountEd/pkg/storage/adapters/framework/out"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	accountTypeCollectionName = []byte("account_type")
)

type logger func(level, msg string)
type mongoactions *out.MongoActions

type Adapter struct {
	mongoActions      mongoactions
	log               logger
	ResponseItemLimit int32
}

func NewAdapter(m mongoactions, l logger) *Adapter {
	return &Adapter{
		mongoActions: m,
		log:          l,
	}
}

func (a *Adapter) InsertAccountType(data *[]byte) (*[]byte, error) {
	var in map[string]interface{}
	var output []byte

	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))
	t := a.mongoActions.StorageAdapter.GetMongoActions().GetTimeStamp()

	e := json.Unmarshal(*data, &in)
	if e != nil {
		a.log("error", fmt.Sprintf("insert account type failed: %v", e))
		return &output, e
	}

	in["created_at"] = t
	in["modified_at"] = t

	result, e := collection.InsertOne(context.TODO(), bson.M(in))

	if e != nil {
		a.log("error", e.Error())

		if mongoError, ok := e.(mongo.WriteException); ok {
			a.mongoActions.WriteError.WriteErrorMsg = &mongoError.WriteErrors
			return nil, errors.New(a.mongoActions.WriteError.Error())
		}

		return nil, e
	}

	id := storage.InsertId(result.InsertedID.(primitive.ObjectID).Hex())
	r := storage.MongoInsertOutput{InsertId: &id, TimeStamp: &t}

	output, _ = json.Marshal(&r)

	return &output, nil
}

func (a *Adapter) RemoveAccountType(id *string) (*[]byte, error) {
	var s map[string]interface{}

	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))
	mid, e := primitive.ObjectIDFromHex(*id)

	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	e = collection.FindOneAndDelete(context.TODO(), bson.D{{Key: "_id", Value: mid}}, nil).Decode(&s)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	v, e := json.Marshal(&s)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	return &v, nil
}

func (a *Adapter) UpdateAccountType(name, id *string) (*int64, error) {
	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))

	t := a.mongoActions.StorageAdapter.GetMongoActions().GetTimeStamp()

	objectId, e := primitive.ObjectIDFromHex(*id)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	d := bson.D{{Key: "$set", Value: bson.D{{Key: "account_type", Value: name}, {Key: "modified_at", Value: t}}}}

	r, e := collection.UpdateOne(context.TODO(), filter, d)
	if e != nil {
		a.log("error", e.Error())

		if mongoError, ok := e.(mongo.WriteException); ok {
			a.mongoActions.WriteError.WriteErrorMsg = &mongoError.WriteErrors
			return nil, errors.New(a.mongoActions.WriteError.Error())
		}

		return nil, e
	}

	return &r.MatchedCount, nil
}
