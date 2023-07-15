package server

import (
	"fmt"
	"os"
	"reflect"

	"golang.org/x/exp/slog"
)

type environment struct {
	AccessKey                string
	DynamoEndpoint           string
	DynamoTableName          string
	Port                     string
	PreRegistrationParameter string
	Region                   string
	RoleToAssume             string
	SecretAccessKey          string
}

type Adapter struct {
	log *slog.Logger
}

func NewAdapter(log *slog.Logger) *Adapter {
	return &Adapter{
		log: log,
	}
}

// LoadSessionConfig retrieves the AWS configuration from environment variables, validates the required values, and returns the configuration in
// a map. If a mandatory configuration value is missing, it logs an error message and triggers the program to exit with an error code.
func (a *Adapter) LoadSessionConfig() *map[string]interface{} {
	var out = make(map[string]interface{})
	var s string

	env := environment{
		AccessKey:                os.Getenv("AWS_ACCESS_KEY_ID"),
		DynamoEndpoint:           os.Getenv("DYNAMODB_ENDPOINT"),
		DynamoTableName:          os.Getenv("DYNAMODB_TABLE_NAME"),
		Port:                     os.Getenv("PORT"),
		PreRegistrationParameter: os.Getenv("AWS_PRE_REGISTRATION_PARAM"),
		Region:                   os.Getenv("AWS_REGION"),
		RoleToAssume:             os.Getenv("AWS_ROLE_TO_ASSUME"),
		SecretAccessKey:          os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	val := reflect.ValueOf(&env).Elem()

	for i := 0; i < val.NumField(); i++ {
		out[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if k == "Port" {
				continue
			}

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
