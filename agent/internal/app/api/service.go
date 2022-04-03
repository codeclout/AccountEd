package api

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type ServiceConfig struct {
	Service     SConf             `hcl:"service,block"`
	HealthCheck HealthCheckConfig `hcl:"healthcheck,block"`
}

type HealthCheckConfig struct {
	Interval string `hcl:"interval"`
	Retries  int16  `hcl:"retries"`
	Timeout  string `hcl:"timeout"`
}

type SConf struct { //nolint:maligned
	CacheDriver              string `hcl:"cache_driver"`
	DatabaseDriver           string `hcl:"db_driver"`
	DatabaseConnectionString string `hcl:"db_connection_string"`
	HostName                 string `hcl:"address"`
	Port                     int16  `hcl:"port"`
	Protocol                 string `hcl:"protocol,label"`
	UseCache                 bool   `hcl:"use_cache"`
}

func (c ServiceConfig) GetConfig() (*ServiceConfig, error) {
	wd, wde := os.Getwd()
	if wde != nil {
		return nil, wde
	}

	configFileLocation := filepath.Join(wd, "service.hcl")
	e := hclsimple.DecodeFile(configFileLocation, nil, &c)

	if e != nil {
		return nil, e
	}

	return &c, nil
}
