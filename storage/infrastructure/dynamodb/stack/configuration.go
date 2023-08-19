package stack

import (
	"fmt"
	"os"
	"reflect"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type environment struct {
	AccessKey       string
	AWSRegion       string
	RoleToAssume    string
	SecretAccessKey string
	SessionLabel    string
}

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		monitor: monitor,
	}
}

func (a *Adapter) LoadStorageInfrastructureConfig(base map[string]interface{}) *map[string]interface{} {
	var override = make(map[string]interface{})
	var s string

	envConfig := environment{
		AccessKey:       os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSRegion:       os.Getenv("AWS_REGION"),
		RoleToAssume:    os.Getenv("AWS_ROLE_TO_ASSUME"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionLabel:    os.Getenv("AWS_SESSION_LABEL"),
	}

	val := reflect.ValueOf(&envConfig).Elem()

	for i := 0; i < val.NumField(); i++ {
		override[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	for k, v := range override {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.monitor.LogGenericError(fmt.Sprintf("Storage Infrastructure:%s is not defined in the environment", k))
				os.Exit(1)
			}
		default:
			panic("invalid Notification configuration type")
		}
	}

	for k, v := range base {
		if k == "SLARoutes" {
			continue
		}

		override[k] = v
	}

	return &override
}
