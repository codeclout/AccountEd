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

type Config struct {
	DbConnectionTimeout int64  `hcl:"database_connection_timeout"`
	DbName              string `hcl:"database_name"`
	DefaultListLimit    int64  `hcl:"default_list_count_limit"`
}

type RequestConfig struct {
}

type Adapter struct {
	config Config
	log    l
}

func NewAdapter(logger l) *Adapter {
	return &Adapter{
		log: logger,
	}
}

func (a *Adapter) GetConfig(path []byte) []byte {
	wd, _ := os.Getwd()

	configFileLocation := filepath.Join(wd, string(path))
	e := hclsimple.DecodeFile(configFileLocation, nil, &a.config)

	if e != nil {
		x, ok := e.(hcl.Diagnostics)

		if ok {
			a.log("fatal", fmt.Sprintf("Failed to load runtime config: %s", e.(hcl.Diagnostics)[0].Summary))
		} else {
			a.log("fatal", fmt.Sprintf("Failed to get runtime config: %v", x))
		}
	}

	b, e := json.Marshal(a.config)

	if e != nil {
		a.log("fatal", fmt.Sprintf("unmarshall error: %v", e))
	}

	return b
}
