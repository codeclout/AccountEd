package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	ports "github.com/codeclout/AccountEd/ports/framework/out/db"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *Adapter) InsertAccountType(data []byte) (ports.InsertID, error) {
	var in map[string]interface{}

	collection := a.db.Collection(string(accountTypeCollectionName))

	e := json.Unmarshal(data, &in)
	if e != nil {
		a.log("error", fmt.Sprintf("invalid payload: %v", e))
		return ports.InsertID{}, e
	}

	result, e := collection.InsertOne(a.ctx, bson.M(in))

	if e != nil {
		a.log("error", fmt.Sprintf("error inserting account type: %v", e))
		return ports.InsertID{}, e
	}

	return ports.InsertID{InsertedID: result.InsertedID}, nil
}

func (a *Adapter) GetAccountTypes(v int64) ([]byte, error) {
	// slice of map
	var temp []bson.M

	collection := a.db.Collection(string(accountTypeCollectionName))
	limit, ok := a.config["DefaultListLimit"].(float64)

	if !ok {
		a.log("error", fmt.Sprintf("Expecting float64 and received %T", limit))
		return []byte{}, errors.New("invalid type for limit")
	}

	if v != -1 {
		limit = float64(v)
	}

	o := options.Find().SetLimit(int64(limit))
	cs, e := collection.Find(a.ctx, bson.M{}, o)

	if e != nil {
		a.log("error", fmt.Sprintf("error inserting account type: %v", e))
		return []byte{}, e
	}

	if e = cs.All(a.ctx, &temp); e != nil {
		a.log("error", fmt.Sprintf("Results decode failed: %v", e))
		return []byte{}, e
	}

	b, _ := json.Marshal(temp)

	return b, nil
}
