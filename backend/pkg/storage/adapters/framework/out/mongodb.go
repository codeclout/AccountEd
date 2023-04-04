package out

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoActions struct {
	BulkWriteError MongoBulkWriteError
	Cancel         context.CancelFunc
	Client         *mongo.Client
	Db             *mongo.Database
	WriteError     MongoWriteError
	StorageAdapter *Adapter
}

type MongoWriteError struct {
	WriteErrorMsg *mongo.WriteErrors
}

type MongoBulkWriteError struct {
	BulkWriteErrorExceptionMsg *mongo.BulkWriteException
}

func (e MongoBulkWriteError) Error() (s string) {
	for i, v := range e.BulkWriteErrorExceptionMsg.WriteErrors {
		if i == 0 {
			s = v.Message
			continue
		}

		s = s + "," + v.Message
	}

	return s
}

func (e MongoWriteError) Error() string {
	var s string

	for i, v := range *e.WriteErrorMsg {
		if i == 0 {
			s = v.Message
			continue
		}

		s = s + "," + v.Message
	}

	return s
}

func (a *Adapter) NewMongoDb(srv string) *MongoActions {

	s, ok := a.RuntimeConfig["DbConnectionTimeout"].(float64)
	if !ok {
		a.log("fatal", fmt.Sprintf("Expecting float64 and received %T", s))
	}

	dbname, k := a.RuntimeConfig["DbName"].(string)
	if !k && len(dbname) >= 3 {
		a.log("fatal", fmt.Sprintf("Expecting database name with 3 or greater characters and recevied %s", dbname))
	}

	t := time.Duration(int64(s)) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)

	client, e := mongo.Connect(ctx, options.Client().ApplyURI(srv))
	if e != nil {
		a.log("fatal", fmt.Sprintf("db connection failed %v", e))
	}

	database := client.Database(dbname)

	return &MongoActions{
		Cancel:         cancel,
		Client:         client,
		Db:             database,
		StorageAdapter: a,
	}
}

func (m *MongoActions) GetTimeStamp() primitive.DateTime {
	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()

	return primitive.NewDateTimeFromTime(t)
}

func (m *MongoActions) Initialize() {
	e := m.Client.Ping(context.TODO(), readpref.Primary())
	if e != nil {
		m.StorageAdapter.log("fatal", fmt.Sprintf("db ping failed: %v", e))
	}

	m.StorageAdapter.log("info", "successfully connected to the database")
}

func (m *MongoActions) CloseConnection() {
	defer m.Cancel()
	if e := m.Client.Disconnect(context.TODO()); e != nil {
		panic(e)
	}
}
