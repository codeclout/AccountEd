package server

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	membertypes "github.com/codeclout/AccountEd/members/member-types"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"golang.org/x/exp/slog"
)

type environment struct {
	AWSRegion                string
	AWSRolePreRegistration   string
	PreRegistrationParameter string
}

type server struct {
	GetOnlyConstraint bool    `hcl:"is_app_get_only" json:"is_app_get_only"`
	Name              string  `hcl:"application_name" json:"application_name"`
	RoutePrefix       string  `hcl:"route_prefix" json:"route_prefix"`
	SLARoutes         float64 `hcl:"sla_routes" json:"sla_routes"`
}

type Adapter struct {
	log                     *slog.Logger
	staticConfigurationPath membertypes.ConfigurationPath
}

func NewAdapter(log *slog.Logger, configPath membertypes.ConfigurationPath) *Adapter {
	return &Adapter{
		log:                     log,
		staticConfigurationPath: configPath,
	}
}

// LoadMemberConfig reads and decodes the configuration file located by the "path" constant, and
// populates a server instance. The server instance is then converted to a JSON object which is
// unmarshalled into a map[string]interface{}. This method returns a pointer to the map.
// In case of any errors, it logs the error message and panics with the corresponding error.
func (a *Adapter) LoadMemberConfig() *map[string]interface{} {
	var configuration server
	var out = make(map[string]interface{})
	var s string

	workingDirectory, _ := os.Getwd()
	fileLocation := filepath.Join(workingDirectory, string(a.staticConfigurationPath))

	e := hclsimple.DecodeFile(fileLocation, nil, &configuration)
	if e != nil {
		var x hcl.Diagnostics
		ok := errors.Is(e, x)

		if ok {
			a.log.Error("fatal", fmt.Sprintf("Failed to load runtime staticConfig: %s", e.(hcl.Diagnostics)[0].Summary))
			panic(e)
		} else {
			a.log.Error("fatal", fmt.Sprintf("Failed to get runtime staticConfig: %v", x))
			panic(e)
		}
	}

	env := environment{
		AWSRegion:                os.Getenv("AWS_REGION"),
		AWSRolePreRegistration:   os.Getenv("AWS_PRE_REGISTRATION_ROLE"),
		PreRegistrationParameter: os.Getenv("AWS_PRE_REGISTRATION_HASH_PARAM"),
	}

	sval := reflect.ValueOf(&env).Elem()
	for i := 0; i < sval.NumField(); i++ {
		out[sval.Type().Field(i).Name] = sval.Field(i).Interface()
	}

	val := reflect.ValueOf(&configuration).Elem()
	for i := 0; i < val.NumField(); i++ {
		out[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	for k, v := range out {
		switch x := v.(type) {
		case string:
			if x == (s) {
				a.log.Error(fmt.Sprintf("Members:%s is not defined in the environment", k))
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
