package server

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"

	membertypes "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type environment struct {
	AWSRegion                string
	AWSRolePreRegistration   string
	Domain                   string
	Port                     string
	PreRegistrationParameter string
	RuntimeEnvironment       string
}

type metadataAndSettings struct {
	Metadata Metadata `hcl:"Metadata,block"`
	Settings Settings `hcl:"Settings,block"`
}

type Metadata struct {
	GetOnlyConstraint bool   `hcl:"is_app_get_only"`
	Name              string `hcl:"application_name"`
	Version           string `hcl:"version"`
}

type Settings struct {
	SLARoutes float64 `hcl:"sla_routes"`
}

type Adapter struct {
	monitor                 monitoring.Adapter
	staticConfigurationPath membertypes.ConfigurationPath
}

func NewAdapter(monitor monitoring.Adapter, configPath membertypes.ConfigurationPath) *Adapter {
	return &Adapter{
		monitor:                 monitor,
		staticConfigurationPath: configPath,
	}
}

//nolint:funlen
func (a *Adapter) LoadMemberConfig() *map[string]interface{} {
	var metadataAndSettings metadataAndSettings
	var out = make(map[string]interface{})
	var s string

	workingDirectory, _ := os.Getwd()
	fileLocation := filepath.Join(workingDirectory, string(a.staticConfigurationPath))

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
		AWSRegion:                os.Getenv("AWS_REGION"),
		AWSRolePreRegistration:   os.Getenv("AWS_PRE_REGISTRATION_ROLE"),
		Domain:                   os.Getenv("DOMAIN"),
		Port:                     os.Getenv("PORT"),
		PreRegistrationParameter: os.Getenv("AWS_PRE_REGISTRATION_HASH_PARAM"),
		RuntimeEnvironment:       os.Getenv("ENVIRONMENT"),
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
				a.monitor.LogGenericError(fmt.Sprintf("Members:%s is not defined in the environment", k))
				os.Exit(1)
			}
		case bool:
			continue
		case float64:
			continue
		default:
			panic("invalid Members configuration type")
		}
	}

	return &out
}
