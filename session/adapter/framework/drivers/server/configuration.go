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
	AccessKey                string
	Port                     string
	PreRegistrationParameter string
	Region                   string
	RoleToAssume             string
	SecretAccessKey          string
	SessionTableName         string
}

type metadataAndSettings struct {
	Metadata Metadata `hcl:"Metadata,block"`
	Settings Settings `hcl:"Settings,block"`
}

type Metadata struct {
	ServiceName string `hcl:"service"`
	Version     string `hcl:"version"`
}

type Settings struct {
	MaxGRPCResponseTime float64 `hcl:"max_grpc_response_time"`
}

type Adapter struct {
	monitor                 monitoring.Adapter
	staticConfigurationPath string
}

func NewAdapter(configurationPath string, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		monitor:                 monitor,
		staticConfigurationPath: configurationPath,
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
					a.monitor.LogGenericError(fmt.Sprintf("Failed to load member runtime staticConfig: %s", x))
				}
			}
			panic(e)
		}
	}

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
			if k == "Port" {
				continue
			}

			if x == (s) {
				a.monitor.LogGenericError(fmt.Sprintf("AWS:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case float64:
			continue
		default:
			panic("invalid AWS configuration type")
		}
	}

	return &out
}
