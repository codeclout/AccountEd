package service_config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type l func(level string, msg string)

// DatabaseConfig represents the static runtime configuration of the service
type DatabaseConfig struct {
	StaticDbConfig StaticDbConfig `hcl:"db,block"`
}

type StaticDbConfig struct {
	DbConnectionTimeout int64  `hcl:"database_connection_timeout"`
	DbName              string `hcl:"database_name"`
	DefaultListLimit    int64  `hcl:"default_list_count_limit"`
	UseDynamoDb         bool   `hcl:"use_dynamo_db"`
	UseMongoDb          bool   `hcl:"use_mongo_db"`
}

type ENV struct {
	AwsAccessKey       string
	AwsRegion          string
	AwsRoleToAssume    string
	AwsSecretAccessKey string
	AwsSessionName     string
	DbConnectionParam  string
	DynamoDbEndpoint   string
	DynamoDbTableName  string
	MapKey             string
}

type RuntimeConfig struct {
	*StaticDbConfig
	*ENV
}

type RequestConfig struct {
}

type Adapter struct {
	config  DatabaseConfig
	log     l
	runtime RuntimeConfig
}

func NewAdapter(logger l) *Adapter {
	return &Adapter{
		log: logger,
	}
}

func (a *Adapter) GetConfig(path []byte) []byte {
	var dbStaticConfig DatabaseConfig
	wd, _ := os.Getwd()

	configFileLocation := filepath.Join(wd, string(path))
	e := hclsimple.DecodeFile(configFileLocation, nil, &dbStaticConfig)

	if e != nil {
		x, ok := e.(hcl.Diagnostics)

		if ok {
			a.log("fatal", fmt.Sprintf("Failed to load runtime dbStaticConfig: %s", e.(hcl.Diagnostics)[0].Summary))
		} else {
			a.log("fatal", fmt.Sprintf("Failed to get runtime dbStaticConfig: %v", x))
		}
	}

	env := ENV{
		AwsAccessKey:       os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsRoleToAssume:    os.Getenv("AWS_ROLE_TO_ASSUME"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsSessionName:     os.Getenv("AWS_SESSION_NAME"),
		DbConnectionParam:  os.Getenv("DB_CONNECTION_PARAM"),
		DynamoDbEndpoint:   os.Getenv("DYNAMODB_ENDPOINT"),
		DynamoDbTableName:  os.Getenv("DYNAMODB_TABLE_NAME"),
		MapKey:             os.Getenv("GCP_MAP_API_KEY"),
	}

	a.runtime.StaticDbConfig = &dbStaticConfig.StaticDbConfig
	a.runtime.ENV = &env

	b, e := json.Marshal(a.runtime)

	if e != nil {
		a.log("fatal", fmt.Sprintf("json encoding error: %v", e))
	}

	return b
}
