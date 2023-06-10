package server

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
)

type AmazonCloud struct{}

type Adapter struct {
	AmazonCloud
	log *slog.Logger
}

type aws struct {
	AccessKey       string
	DynamoEndpoint  string
	DynamoTableName string
	Region          string
	RoleToAssume    string
	SecretAccessKey string
}

func NewAdapter(log *slog.Logger) *Adapter {
	return &Adapter{
		log: log,
	}
}

// LoadSessionConfig retrieves the AWS configuration from environment variables, validates the required values, and returns the configuration in
// a map. If a mandatory configuration value is missing, it logs an error message and triggers the program to exit with an error code.
func (a *Adapter) LoadSessionConfig() *map[string]interface{} {
	var out map[string]interface{}
	var s string

	awsconfig := aws{
		AccessKey:       os.Getenv("AWS_ACCESS_KEY_ID"),
		DynamoEndpoint:  os.Getenv("DYNAMODB_ENDPOINT"),
		DynamoTableName: os.Getenv("DYNAMODB_TABLE_NAME"),
		Region:          os.Getenv("AWS_REGION"),
		RoleToAssume:    os.Getenv("AWS_ROLE_TO_ASSUME"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	awsenv, _ := json.Marshal(awsconfig)
	_ = json.Unmarshal(awsenv, &out)

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.log.Error(fmt.Sprintf("AWS:%s is not defined in the environment", k))
				os.Exit(1)
			}
		default:
			panic("invalid AWS configuration type")
		}
	}

	return &out
}
