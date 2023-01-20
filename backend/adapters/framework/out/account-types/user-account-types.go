package account_types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/codeclout/AccountEd/adapters/framework/out/storage"
	ports "github.com/codeclout/AccountEd/ports/framework/out/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	accountTypeCollectionName = []byte("account_type")
)

type Adapter struct {
	actions           *storage.MongoActions
	log               func(level, msg string)
	ResponseItemLimit int32
}

func NewAdapter(m *storage.MongoActions, logger func(level, msg string)) *Adapter {
	return &Adapter{
		actions: m,
		log:     logger,
	}
}

func (a *Adapter) InsertAccountType(data *[]byte) (*[]byte, error) {
	var in map[string]interface{}
	var out []byte

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))
	t := a.actions.GetTimeStamp()

	e := json.Unmarshal(*data, &in)
	if e != nil {
		a.log("error", fmt.Sprintf("insert account type failed: %v", e))
		return &out, e
	}

	in["created_at"] = t
	in["modified_at"] = t

	result, e := collection.InsertOne(context.TODO(), bson.M(in))

	if e != nil {
		a.log("error", e.Error())

		if mongoError, ok := e.(mongo.WriteException); ok {
			a.actions.WriteError.Msg = &mongoError.WriteErrors
			return nil, errors.New(a.actions.WriteError.Error())
		}

		return nil, e
	}

	id := ports.InsertId(result.InsertedID.(primitive.ObjectID).Hex())
	r := ports.InsertAccountOutput{InsertId: &id, TimeStamp: &t}

	out, _ = json.Marshal(&r)

	return &out, nil
}

func (a *Adapter) GetAccountTypes(inlimit *int16) (*[]byte, error) {
	// slice of map
	var t []bson.M

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))
	lmt, ok := a.actions.StorageAdapter.RuntimeConfig["DefaultListLimit"].(float64)

	if !ok {
		a.log("error", fmt.Sprintf("Expecting float64 and received %T", lmt))
		return nil, errors.New("invalid type for limit")
	}

	if *inlimit >= 1 && float64(*inlimit) <= lmt {
		lmt = float64(*inlimit)
	}

	o := options.Find().SetLimit(int64(lmt))
	csr, e := collection.Find(context.TODO(), bson.M{}, o)

	if e != nil {
		a.log("error", fmt.Sprintf("error inserting account type: %v", e))
		return nil, e
	}

	if e = csr.All(context.TODO(), &t); e != nil {
		a.log("error", fmt.Sprintf("Results decode failed: %v", e))
		return nil, e
	}

	b, _ := json.Marshal(&t)

	return &b, nil
}

func (a *Adapter) RemoveAccountType(id *string) (*[]byte, error) {
	var s map[string]interface{}

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))
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
	collection := a.actions.Db.Collection(string(accountTypeCollectionName))

	t := a.actions.GetTimeStamp()

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
			a.actions.WriteError.Msg = &mongoError.WriteErrors
			return nil, errors.New(a.actions.WriteError.Error())
		}

		return nil, e
	}

	return &r.MatchedCount, nil
}

func (a *Adapter) GetAccountTypeById(id *string) (*[]byte, error) {
	var (
		e error
		n map[string]interface{}
	)

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))

	objectId, e := primitive.ObjectIDFromHex(*id)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	e = collection.FindOne(context.TODO(), filter).Decode(&n)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	b, e := json.Marshal(&n)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	return &b, e
}
