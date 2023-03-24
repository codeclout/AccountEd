package account_types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/codeclout/AccountEd/storage/adapters/framework/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	accountTypeCollectionName = []byte("account_type")
)

type Adapter struct {
	mongoActions      *out.MongoActions
	log               func(level, msg string)
	ResponseItemLimit int32
}

func NewAdapter(m *out.MongoActions, logger func(level, msg string)) *Adapter {
	return &Adapter{
		mongoActions: m,
		log:          logger,
	}
}

func (a *Adapter) GetAccountTypes(inlimit *int16) (*[]byte, error) {
	// slice of map
	var t []bson.M

	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))
	lmt, ok := a.mongoActions.StorageAdapter.RuntimeConfig["DefaultListLimit"].(float64)

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

func (a *Adapter) GetAccountTypeById(id *string) (*[]byte, error) {
	var (
		e error
		n map[string]interface{}
	)

	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))

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
