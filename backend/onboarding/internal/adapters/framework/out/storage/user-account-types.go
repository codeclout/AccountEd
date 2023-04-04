package storage

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

var (
	accountTypeCollectionName = []byte("account_type")
)

func (a *Adapter) GetAccountTypes(ctx context.Context, inlimit int16) (*[]internal.AccountTypeOut, error) {
	// slice of map
	var out []internal.AccountTypeOut

	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))
	defaultLimit, ok := a.mongoActions.StorageAdapter.RuntimeConfig["DefaultListLimit"].(float64)

	if !ok {
		a.log("error", fmt.Sprintf("invalid GetAccountTypes default limit %T", defaultLimit))
		return nil, errors.New("invalid Default Limit - GetAccountTypes")
	}

	if inlimit >= 1 && float64(inlimit) <= defaultLimit {
		defaultLimit = float64(inlimit)
	}

	o := options.Find().SetLimit(int64(defaultLimit))
	csr, e := collection.Find(ctx, bson.M{}, o)

	if e != nil {
		a.log("error", fmt.Sprintf("error inserting account type: %v", e))
		return nil, e
	}

	if e = csr.All(ctx, &out); e != nil {
		a.log("error", fmt.Sprintf("Results decode failed: %v", e))
		return nil, e
	}

	return &out, nil
}

func (a *Adapter) GetAccountTypeById(ctx context.Context, in internal.AccountTypeIn) (*internal.AccountTypeOut, error) {
	var out internal.AccountTypeOut

	collection := a.mongoActions.Db.Collection(string(accountTypeCollectionName))

	objectId, e := primitive.ObjectIDFromHex(in.Id)
	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	e = collection.FindOne(ctx, filter).Decode(&out)
	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	return &out, e
}
