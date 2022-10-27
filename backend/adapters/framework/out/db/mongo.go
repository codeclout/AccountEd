package db

import (
	"context"
	"fmt"
	"time"

	ports "github.com/codeclout/AccountEd/ports/framework/out/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Adapter struct {
	cancel context.CancelFunc
	client *mongo.Client
	ctx    context.Context
	logger func(l string, m string)
}

func NewAdapter(timeout int, logger func(level string, msg string), uri string) (*Adapter, error) {
	t := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)

	client, e := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if e != nil {
		logger("fatal", fmt.Sprintf("db connection failed %v", e))
	}

	e = client.Ping(ctx, readpref.Primary())
	if e != nil {
		logger("fatal", fmt.Sprintf("db ping failed: %v", e))
	}

	a := Adapter{
		cancel: cancel,
		client: client,
		ctx:    ctx,
		logger: logger,
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

func (a *Adapter) InsertAccountType(coll string, data string) (ports.InsertID, error) {
	database := a.client.Database("accountEd")
	collection := database.Collection(coll)

	result, e := collection.InsertOne(context.TODO(), bson.M{"account_type": data})

	if e != nil {
		fmt.Printf("error inserting account type: %v", e)
		return ports.InsertID{}, e
	}

	return ports.InsertID{InsertedID: result.InsertedID}, nil
}
