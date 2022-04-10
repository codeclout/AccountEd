package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Adapter struct {
	client  *mongo.Client
	ctx     context.Context
	db      string
	timeout time.Duration
}

func NewAdapter(url, db string, timeout int) (*Adapter, error) {
	t := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	client, e := mongo.Connect(ctx, options.Client().ApplyURI(db))
	if e != nil {
		log.Fatalf("db connection failed: %v", e)
	}

	e = client.Ping(ctx, readpref.Primary())
	if e != nil {
		log.Fatalf("db ping failed: %v", e)
	}

	a := &Adapter{
		client:  client,
		ctx:     ctx,
		db:      "",
		timeout: time.Duration(t) * time.Second,
	}

	defer a.CloseConnection()

	return a, nil
}

func (a *Adapter) CloseConnection() {
	if e := a.client.Disconnect(a.ctx); e != nil {
		panic(e)
	}
}

func (a *Adapter) LogDataInteraction(k, v string) {}
