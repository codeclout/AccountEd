package hcl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type l func(level string, msg string)

// Config represents the runtime configuration of the service
type Config struct {
	DbConnectionTimeout int64  `hcl:"database_connection_timeout"`
	DbName              string `hcl:"database_name"`
	DefaultListLimit    int64  `hcl:"default_list_count_limit"`
}

type ENV struct {
	AwsAccessKey       string
	AwsRegion          string
	AwsRoleToAssume    string
	AwsSecretAccessKey string
}

type RuntimeConfig struct {
	*Config
	*ENV
}

type RequestConfig struct {
}

type Adapter struct {
	config  Config
	log     l
	runtime RuntimeConfig
}

func NewAdapter(logger l) *Adapter {
	return &Adapter{
		log: logger,
	}
}

func (a *Adapter) GetConfig(path []byte) []byte {
	var config Config
	wd, _ := os.Getwd()

	configFileLocation := filepath.Join(wd, string(path))
	e := hclsimple.DecodeFile(configFileLocation, nil, &config)

	if e != nil {
		x, ok := e.(hcl.Diagnostics)

		if ok {
			a.log("fatal", fmt.Sprintf("Failed to load runtime config: %s", e.(hcl.Diagnostics)[0].Summary))
		} else {
			a.log("fatal", fmt.Sprintf("Failed to get runtime config: %v", x))
		}
	}

	env := ENV{
		AwsAccessKey:       os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsRoleToAssume:    os.Getenv("AWS_ROLE_TO_ASSUME"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	a.runtime.Config = &config
	a.runtime.ENV = &env

	b, e := json.Marshal(a.runtime)

	if e != nil {
		a.log("fatal", fmt.Sprintf("json encoding error: %v", e))
	}

	return b
}
