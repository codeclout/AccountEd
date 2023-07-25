package server

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type environment struct {
	DynamoEndpoint  string
	DynamoTableName string
	Port            string
	Region          string
	RoleToAssume    string
}

type server struct {
	SLARoutes float64 `hcl:"sla_routes" json:"sla_routes"`
}

type Adapter struct {
	monitor                 monitoring.Adapter
	staticConfigurationPath string
}

func NewAdapter(monitor monitoring.Adapter, staticConfigPath string) *Adapter {
	return &Adapter{
		monitor:                 monitor,
		staticConfigurationPath: staticConfigPath,
	}
}

func (a *Adapter) LoadSessionConfig() *map[string]interface{} {
	var configuration server
	var out = make(map[string]interface{})
	var s string

	workingDirectory, _ := os.Getwd()
	fileLocation := filepath.Join(workingDirectory, a.staticConfigurationPath)

	e := hclsimple.DecodeFile(fileLocation, nil, &configuration)
	if e != nil {
		var hclError hcl.Diagnostics
		ok := errors.Is(e, hclError)

		if ok {
			a.monitor.LogGenericError(fmt.Sprintf("Failed to load runtime staticConfig: %s", e.(hcl.Diagnostics)[0].Summary))
			panic(e)
		} else {
			a.monitor.LogGenericError(fmt.Sprintf("Failed to get runtime staticConfig: %v", hclError))
			panic(e)
		}
	}

	env := environment{
		DynamoEndpoint:  os.Getenv("DYNAMODB_ENDPOINT"),
		DynamoTableName: os.Getenv("DYNAMODB_TABLE_NAME"),
		Port:            os.Getenv("PORT"),
		Region:          os.Getenv("AWS_REGION"),
		RoleToAssume:    os.Getenv("AWS_ROLE_TO_ASSUME"),
	}

	val := reflect.ValueOf(&env).Elem()
	for i := 0; i < val.NumField(); i++ {
		out[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	staticConfig := reflect.ValueOf(&configuration).Elem()
	for i := 0; i < staticConfig.NumField(); i++ {
		out[staticConfig.Type().Field(i).Name] = staticConfig.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.monitor.LogGenericError(fmt.Sprintf("Storage:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case float64:
			continue
		default:
			panic("invalid Storage configuration type")
		}
	}

	return &out
}
