package out

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type logger func(level, msg string)
type runtimeConfig map[string]interface{}

type Adapter struct {
	cloudSession            *aws.Config
	ctx                     context.Context
	dynamoActions           *DynamoActions
	isRegisteredDynamo      bool
	isRegisteredMongo       bool
	log                     logger
	mongoActions            *MongoActions
	mongodbConnectionString *string
	RuntimeConfig           runtimeConfig
}

func handleError(e error, logger logger) {
	if e != nil {
		logger("fatal", e.Error())
	}
}

func NewAdapter(config []byte, logger logger, mongoConnection *string, cloudConfig *aws.Config) (*Adapter, error) {
	var cfg runtimeConfig

	e := json.Unmarshal(config, &cfg)
	handleError(e, logger)

	a := Adapter{
		cloudSession:            cloudConfig,
		ctx:                     context.TODO(),
		isRegisteredDynamo:      false,
		isRegisteredMongo:       false,
		log:                     logger,
		mongodbConnectionString: mongoConnection,
		RuntimeConfig:           cfg,
	}

	return &a, nil
}

func (a *Adapter) Initialize() {
	if a.RuntimeConfig["UseMongoDb"].(bool) && !a.isRegisteredMongo {
		mongoactions := a.NewMongoDb(*a.mongodbConnectionString)

		mongoactions.Initialize()
		a.isRegisteredMongo = true
		a.mongoActions = mongoactions
	}

	if a.RuntimeConfig["UseDynamoDb"].(bool) && !a.isRegisteredDynamo {
		dynamoactions := a.NewDynamoDb(a.cloudSession)

		dynamoactions.Initialize()
		a.isRegisteredDynamo = true

		a.dynamoActions = dynamoactions
	}
}

func (a *Adapter) CloseConnection() {
	if a.RuntimeConfig["UseMongoDb"].(bool) && a.isRegisteredMongo {
		a.mongoActions.CloseConnection()
	}
}

func (a *Adapter) GetMongoActions() *MongoActions {
	return a.mongoActions
}
