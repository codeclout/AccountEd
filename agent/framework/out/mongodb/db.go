package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Adapter struct {
	client *mongo.Client
	ctx    context.Context
}

func NewAdapter(uri string, timeout int) (*Adapter, error) {
	t := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	client, e := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if e != nil {
		log.Panicf("db connection failed: %v", e)
		return &Adapter{}, e
	}

	e = client.Ping(ctx, readpref.Primary())
	if e != nil {
		log.Panicf("db ping failed: %v", e)
		return &Adapter{}, e
	}

	a := Adapter{
		client: client,
		ctx:    ctx,
	}

	defer a.CloseConnection()

	return &a, nil
}

func (a *Adapter) CloseConnection() {
	if e := a.client.Disconnect(a.ctx); e != nil {
		panic(e)
	}
}
