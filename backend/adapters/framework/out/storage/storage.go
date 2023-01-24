package storage

import (
	"context"
	"encoding/json"
)

type logger func(level, msg string)
type runtimeConfig map[string]interface{}

type Adapter struct {
	ctx               context.Context
	log               logger
	isRegisteredMongo bool
	mongoActions      *MongoActions
	RuntimeConfig     runtimeConfig
	uri               *string
}

func handleError(e error, logger logger) {
	if e != nil {
		logger("fatal", e.Error())
	}
}

func NewAdapter(config []byte, logger logger, uri *string) (*Adapter, error) {
	var cfg runtimeConfig

	e := json.Unmarshal(config, &cfg)
	handleError(e, logger)

	a := Adapter{
		ctx:               context.TODO(),
		log:               logger,
		isRegisteredMongo: false,
		RuntimeConfig:     cfg,
		uri:               uri,
	}

	return &a, nil
}

func (a *Adapter) Initialize() {
	if a.RuntimeConfig["UseMongoDb"].(bool) && !a.isRegisteredMongo {
		dbActions := a.NewMongoDb(*a.uri)

		dbActions.Initialize()
		a.isRegisteredMongo = true
		a.mongoActions = dbActions
	}
}

func (a *Adapter) CloseConnection() {
	if a.RuntimeConfig["UseMongoDb"].(bool) && a.isRegisteredMongo {
		a.mongoActions.CloseConnection()
	}
}

func (a *Adapter) GetMongoAccountTypeActions() *MongoActions {
	return a.mongoActions
}
