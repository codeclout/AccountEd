package server

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/pkg/errors"

	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type environment struct {
	AWSRegion    string
	Port         string
	RoleToAssume string
}

type metadataAndSettings struct {
	Metadata Metadata `hcl:"Metadata,block"`
	Settings Settings `hcl:"Settings,block"`
}

type Metadata struct {
	Name     string `hcl:"application_name"`
	DemoMode bool   `hcl:"demo_mode"`
	Version  string `hcl:"version"`
}

type Settings struct {
	DynamoDemo          string  `hcl:"dynamodb_demo" json:"DynamoDemo"`
	DynamoFipsUSEast1   string  `hcl:"dynamodb_fips_us_east_1" json:"dynamoFipsUSEast1"`
	DynamoFipsUSEast2   string  `hcl:"dynamodb_fips_us_east_2" json:"dynamoFipsUSEast2"`
	DynamoFipsUSWest1   string  `hcl:"dynamodb_fips_us_west_1" json:"DynamoFipsUSWest1"`
	DynamoFipsUSWest2   string  `hcl:"dynamodb_fips_us_west_2" json:"DynamoFipsUSWest2"`
	DynamoStreamUSEast1 string  `hcl:"dynamodb_stream_us_east_1" json:"DynamoStreamUSEast1"`
	DynamoStreamUSEast2 string  `hcl:"dynamodb_stream_us_east_2" json:"DynamoStreamUSEast2"`
	DynamoStreamUSWest1 string  `hcl:"dynamodb_stream_us_west_1" json:"DynamoStreamUSWest1"`
	DynamoStreamUSWest2 string  `hcl:"dynamodb_stream_us_west_2" json:"DynamoStreamUSWest2"`
	SLARoutes           float64 `hcl:"sla_routes" json:"sla_routes"`
	UseDynamoWtihStream bool    `hcl:"use_dynamodb_with_stream" json:"UseDynamoWtihStream"`
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

func (a *Adapter) LoadStorageConfig() *map[string]interface{} {
	var metadataAndSettings metadataAndSettings
	var out = make(map[string]interface{})
	var s string

	workingDirectory, _ := os.Getwd()
	fileLocation := filepath.Join(workingDirectory, a.staticConfigurationPath)

	e := hclsimple.DecodeFile(fileLocation, nil, &metadataAndSettings)
	if e != nil {
		var x hcl.Diagnostics
		if errors.As(e, &x) {
			for _, x := range x {
				if x.Severity == hcl.DiagError {
					a.monitor.LogGenericError(fmt.Sprintf("Failed to load storage runtime staticConfig: %s", x))
				}
			}
			panic(e)
		}
	}

	env := environment{
		AWSRegion:    os.Getenv("AWS_REGION"),
		Port:         os.Getenv("PORT"),
		RoleToAssume: os.Getenv("AWS_ROLE_TO_ASSUME"),
	}

	runtimeEnv := reflect.ValueOf(&env).Elem()
	for i := 0; i < runtimeEnv.NumField(); i++ {
		out[runtimeEnv.Type().Field(i).Name] = runtimeEnv.Field(i).Interface()
	}

	metadata := reflect.ValueOf(&metadataAndSettings.Metadata).Elem()
	for i := 0; i < metadata.NumField(); i++ {
		out[metadata.Type().Field(i).Name] = metadata.Field(i).Interface()
	}

	settings := reflect.ValueOf(&metadataAndSettings.Settings).Elem()
	for i := 0; i < settings.NumField(); i++ {
		out[settings.Type().Field(i).Name] = settings.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.monitor.LogGenericError(fmt.Sprintf("Storage:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case bool:
			continue
		case float64:
			continue
		default:
			panic("invalid Storage configuration type")
		}
	}

	return &out
}
