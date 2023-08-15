package server

import (
	"fmt"
	"os"
	"reflect"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type environment struct {
	AccessKey                string
	Port                     string
	PreRegistrationParameter string
	Region                   string
	RoleToAssume             string
	SecretAccessKey          string
	SessionTableName         string
}

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		monitor: monitor,
	}
}

func (a *Adapter) LoadStorageConfig() *map[string]interface{} {
	var out = make(map[string]interface{})
	var s string

	env := environment{
		AccessKey:                os.Getenv("AWS_ACCESS_KEY_ID"),
		Port:                     os.Getenv("PORT"),
		PreRegistrationParameter: os.Getenv("AWS_PRE_REGISTRATION_PARAM"),
		Region:                   os.Getenv("AWS_REGION"),
		RoleToAssume:             os.Getenv("AWS_ROLE_TO_ASSUME"),
		SecretAccessKey:          os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionTableName:         os.Getenv("SESSION_STORAGE_TABLE_NAME"),
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
				a.monitor.LogGenericError(fmt.Sprintf("AWS:%s is not defined in the environment", k))
				os.Exit(1)
			}
		default:
			panic("invalid AWS configuration type")
		}
	}

	return &out
}
