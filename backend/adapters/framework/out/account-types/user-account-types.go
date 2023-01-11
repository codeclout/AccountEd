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
	actions storage.MongoActions
	log     func(level, msg string)
}

func NewAdapter(m *storage.MongoActions, logger func(level, msg string)) *Adapter {
	return &Adapter{
		actions: *m,
		log:     logger,
	}
}

func (a *Adapter) InsertAccountType(data []byte) (*[]byte, error) {
	var in map[string]interface{}
	var out []byte

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))
	t := a.actions.GetTimeStamp()

	e := json.Unmarshal(data, &in)
	if e != nil {
		a.log("error", fmt.Sprintf("invalid payload: %v", e))
		return &out, e
	}

	in["created_at"] = t
	in["modified_at"] = t

	result, e := collection.InsertOne(context.TODO(), bson.M(in))

	if e != nil {
		a.log("error", e.Error())

		mongoError := e.(mongo.WriteException)
		a.actions.WriteError.Msg = &mongoError.WriteErrors

		return &out, errors.New(a.actions.WriteError.Error())
	}

	id := ports.InsertId(result.InsertedID.(primitive.ObjectID).Hex())
	r := ports.InsertAccountOutput{InsertId: &id, TimeStamp: &t}

	out, _ = json.Marshal(&r)

	return &out, nil
}

func (a *Adapter) GetAccountTypes(v int64) ([]byte, error) {
	// slice of map
	var t []bson.M

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))
	limit, ok := a.actions.StorageAdapter.RuntimeConfig["DefaultListLimit"].(float64)

	if !ok {
		a.log("error", fmt.Sprintf("Expecting float64 and received %T", limit))
		return nil, errors.New("invalid type for limit")
	}

	if v != -1 {
		limit = float64(v)
	}

	o := options.Find().SetLimit(int64(limit))
	cs, e := collection.Find(context.TODO(), bson.M{}, o)

	if e != nil {
		a.log("error", fmt.Sprintf("error inserting account type: %v", e))
		return nil, e
	}

	if e = cs.All(context.TODO(), &t); e != nil {
		a.log("error", fmt.Sprintf("Results decode failed: %v", e))
		return nil, e
	}

	b, _ := json.Marshal(&t)

	return b, nil
}

func (a *Adapter) RemoveAccountType(id string) ([]byte, error) {
	var s map[string]interface{}

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))
	mid, e := primitive.ObjectIDFromHex(id)

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

	return v, nil
}

func (a *Adapter) UpdateAccountType(in []byte) ([]byte, error) {
	var s map[string]string

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))

	_ = json.Unmarshal(in, &s)
	t := a.actions.GetTimeStamp()

	x, e := primitive.ObjectIDFromHex(s["id"])
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	k := bson.D{{Key: "_id", Value: x}}
	d := bson.D{{Key: "$set", Value: bson.D{{Key: "account_type", Value: s["accountType"]}, {Key: "modified_at", Value: t}}}}

	r, e := collection.UpdateOne(context.TODO(), k, d)
	if e != nil {
		a.log("error", e.Error())

		mongoError := e.(mongo.WriteException)
		a.actions.WriteError.Msg = &mongoError.WriteErrors

		return nil, errors.New(a.actions.WriteError.Error())
	}

	b, e := json.Marshal(r)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	return b, nil
}

func (a *Adapter) GetAccountTypeById(in []byte) ([]byte, error) {
	var (
		e error
		n map[string]interface{}
		s map[string]string
	)

	collection := a.actions.Db.Collection(string(accountTypeCollectionName))

	e = json.Unmarshal(in, &s)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	x, e := primitive.ObjectIDFromHex(s["id"])
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	f := bson.D{{Key: "_id", Value: x}}
	e = collection.FindOne(context.TODO(), f).Decode(&n)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	b, e := json.Marshal(&n)
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	return b, e
}
