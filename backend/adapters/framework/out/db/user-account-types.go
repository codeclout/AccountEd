package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	ports "github.com/codeclout/AccountEd/ports/framework/out/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type m map[string]interface{}
type sl func(level, msg string)

var (
	accountTypeCollectionName = []byte("account_type")
	config                    m
)

type Adapter struct {
	cancel context.CancelFunc
	client *mongo.Client
	config m
	ctx    context.Context
	db     *mongo.Database
	log    sl
}

type WriteError struct {
	mongo.WriteException
	Message string
}

func NewAdapter(c []byte, logger sl, uri string) (*Adapter, error) {
	_ = json.Unmarshal(c, &config)

	s, ok := config["DbConnectionTimeout"].(float64)
	if !ok {
		logger("error", fmt.Sprintf("Expecting float64 and received %T", s))
		return &Adapter{}, errors.New("invalid type for limit")
	}

	t := time.Duration(int64(s)) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)

	client, e := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if e != nil {
		logger("fatal", fmt.Sprintf("db connection failed %v", e))
	}

	e = client.Ping(ctx, readpref.Primary())
	if e != nil {
		logger("fatal", fmt.Sprintf("db ping failed: %v", e))
	}

	database := client.Database(config["DbName"].(string))

	a := Adapter{
		cancel: cancel,
		client: client,
		config: config,
		ctx:    context.TODO(),
		db:     database,
		log:    logger,
	}

	logger("info", "db connected")
	return &a, nil
}

func (a *Adapter) CloseConnection() {
	defer a.cancel()
	if e := a.client.Disconnect(a.ctx); e != nil {
		panic(e)
	}
}

func (a WriteError) Error() string {
	var x []string

	if len(a.WriteException.WriteErrors) == 1 {
		return a.WriteException.WriteErrors[0].Message
	}

	for i, v := range a.WriteException.WriteErrors {
		x[i] = v.Message
	}

	return strings.Join(x[:], ",")
}

func (a *Adapter) InsertAccountType(data []byte) (ports.InsertID, error) {
	var in map[string]interface{}
	var out ports.InsertID

	collection := a.db.Collection(string(accountTypeCollectionName))
	t := a.getTimeStamp()

	e := json.Unmarshal(data, &in)
	if e != nil {
		a.log("error", fmt.Sprintf("invalid payload: %v", e))
		return out, e
	}

	in["created_at"] = t
	in["modified_at"] = t

	result, e := collection.InsertOne(a.ctx, bson.M(in))

	if e != nil {
		a.log("error", e.Error())

		mongoError := e.(mongo.WriteException)
		return out, WriteError{Message: e.Error(), WriteException: mongoError}
	}

	return ports.InsertID{InsertedID: result.InsertedID}, nil
}

func (a *Adapter) GetAccountTypes(v int64) ([]byte, error) {
	// slice of map
	var t []bson.M

	collection := a.db.Collection(string(accountTypeCollectionName))
	limit, ok := a.config["DefaultListLimit"].(float64)

	if !ok {
		a.log("error", fmt.Sprintf("Expecting float64 and received %T", limit))
		return nil, errors.New("invalid type for limit")
	}

	if v != -1 {
		limit = float64(v)
	}

	o := options.Find().SetLimit(int64(limit))
	cs, e := collection.Find(a.ctx, bson.M{}, o)

	if e != nil {
		a.log("error", fmt.Sprintf("error inserting account type: %v", e))
		return nil, e
	}

	if e = cs.All(a.ctx, &t); e != nil {
		a.log("error", fmt.Sprintf("Results decode failed: %v", e))
		return nil, e
	}

	b, _ := json.Marshal(&t)

	return b, nil
}

func (a *Adapter) RemoveAccountType(id string) ([]byte, error) {
	var s map[string]interface{}

	collection := a.db.Collection(string(accountTypeCollectionName))
	mid, e := primitive.ObjectIDFromHex(id)

	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	e = collection.FindOneAndDelete(a.ctx, bson.D{{Key: "_id", Value: mid}}, nil).Decode(&s)
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

	collection := a.db.Collection(string(accountTypeCollectionName))

	_ = json.Unmarshal(in, &s)
	t := a.getTimeStamp()

	x, e := primitive.ObjectIDFromHex(s["id"])
	if e != nil {
		a.log("error", e.Error())
		return nil, e
	}

	f := bson.D{{Key: "_id", Value: x}}
	u := bson.D{{Key: "$set", Value: bson.D{{Key: "account_type", Value: s["accountType"]}, {Key: "modified_at", Value: t}}}}

	r, e := collection.UpdateOne(a.ctx, f, u)
	if e != nil {
		a.log("error", e.Error())

		mongoError := e.(mongo.WriteException)
		return nil, WriteError{Message: e.Error(), WriteException: mongoError}
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

	collection := a.db.Collection(string(accountTypeCollectionName))

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
	e = collection.FindOne(a.ctx, f).Decode(&n)
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
