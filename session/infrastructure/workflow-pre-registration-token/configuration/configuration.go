package configuration

import (
	"fmt"
	"os"
	"reflect"

	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type environment struct {
	AccessKey                  string
	Environment                string
	Region                     string
	RoleArn                    string
	SecretAccessKey            string
	SessionLabel               string
	TerraformCloudOrganization string
}

func LoadConfig(monitor monitoring.Adapter) *map[string]interface{} {
	var out = make(map[string]interface{})
	var s string

	envConfig := environment{
		AccessKey:                  os.Getenv("AWS_ACCESS_KEY_ID"),
		Environment:                os.Getenv("ENVIRONMENT"),
		Region:                     os.Getenv("AWS_REGION"),
		RoleArn:                    os.Getenv("AWS_ROLE_TO_ASSUME"),
		SecretAccessKey:            os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionLabel:               os.Getenv("AWS_SESSION_LABEL"),
		TerraformCloudOrganization: os.Getenv("TF_CLOUD_ORGANIZATION"),
	}

	rv := reflect.ValueOf(&envConfig).Elem()

	for i := 0; i < rv.NumField(); i++ {
		out[rv.Type().Field(i).Name] = rv.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				monitor.LogGenericError(fmt.Sprintf("Token Generation Workflow:%s is not defined in the environment", k))
				os.Exit(1)
			}
		default:
			panic("invalid Token Generation Workflow configuration")
		}
	}

	return &out
}
