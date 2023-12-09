package configuration

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
	AccessKey           string
	Environment         string
	Region              string
	SecretAccessKey     string
	StackDeploymentRole string
	SessionLabel        string
}

type metadataAndSettings struct {
	Metadata Metadata `hcl:"Metadata,block"`
	Settings Settings `hcl:"Settings,block"`
}

type Metadata struct {
	ContainerWorkspaceName string `hcl:"container_workspace"`
	ServiceName            string `hcl:"service_name"`
	Version                string `hcl:"version"`
}

type Settings struct {
	CloudBackendHost string `hcl:"cloud_backend_host"`
	CloudBackendOrg  string `hcl:"cloud_backend_org"`
}

func LoadConfig(monitor monitoring.Adapter) *map[string]interface{} {
	var metadataAndSettings metadataAndSettings
	var out = make(map[string]interface{})
	var s string

	workingDirectory, _ := os.Getwd()
	file := filepath.Join(workingDirectory, "./config.hcl")

	e := hclsimple.DecodeFile(file, nil, &metadataAndSettings)
	if e != nil {
		var x hcl.Diagnostics
		if errors.As(e, &x) {
			for _, x := range x {
				if x.Severity == hcl.DiagError {
					monitor.LogGenericError(fmt.Sprintf("Failed to load member runtime staticConfig: %s", x))
				}
			}
			panic(e)
		}
	}

	envConfig := environment{
		AccessKey:           os.Getenv("AWS_ACCESS_KEY_ID"),
		Environment:         os.Getenv("ENVIRONMENT"),
		Region:              os.Getenv("AWS_REGION"),
		SecretAccessKey:     os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionLabel:        os.Getenv("AWS_SESSION_LABEL"),
		StackDeploymentRole: os.Getenv("AWS_ROLE_TO_ASSUME"),
	}

	rv := reflect.ValueOf(&envConfig).Elem()
	for i := 0; i < rv.NumField(); i++ {
		out[rv.Type().Field(i).Name] = rv.Field(i).Interface()
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
				monitor.LogGenericError(fmt.Sprintf("Token Generation Workflow:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case float64:
			continue
		default:
			panic("invalid Token Generation Workflow configuration")
		}
	}

	return &out
}
