package hcl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

func (a *Adapter) GetConfig(path []byte) ([]byte, error) {
	wd, e := os.Getwd()
	if e != nil {
		a.log("error", fmt.Sprintf("os errorr: %v", e))
		return []byte{}, e
	}

	configFileLocation := filepath.Join(wd, string(path))
	e = hclsimple.DecodeFile(configFileLocation, nil, &a.config)

	if e != nil {
		a.log("error", fmt.Sprintf("hcl decode error: %v", e))
		return nil, e
	}

	b, e := json.Marshal(a.config)

	if e != nil {
		a.log("error", fmt.Sprintf("unmarshall error: %v", e))
		return []byte{}, e
	}

	return b, nil
}
