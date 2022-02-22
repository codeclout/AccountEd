package main

import (
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

var configFileLocation string = "./config.hcl"

type AccessControlListConfig struct {
	Read  string `hcl:"read"`
	Write string `hcl:"write"`
}

type RuntimeConfig struct {
	Environment string
	HostName    string

	ACL          AccessControlListConfig `hcl:"acls,block"`
	HealthCheck  HealthCheckConfig       `hcl:"healthcheck,block"`
	Organization string                  `hcl:"organization"`
	Service      ServiceConfig           `hcl:"service,block"`
}

type HealthCheckConfig struct {
	Interval string `hcl:"interval"`
	Retries  int16  `hcl:"retries"`
	Timeout  string `hcl:"timeout"`
}

type ServiceConfig struct {
	CacheDriver              string `hcl:"cache_driver"`
	DatabaseDriver           string `hcl:"db_driver"`
	DatabaseConnectionString string `hcl:"db_connection_string"`
	HostName                 string `hcl:"address"`
	Port                     int16  `hcl:"port"`
	Protocol                 string `hcl:"protocol,label"`
}

func getConfig() *RuntimeConfig {
	config := RuntimeConfig{
		Environment: os.Getenv("ENVIRONMENT"),
	}

	e := hclsimple.DecodeFile(configFileLocation, nil, &config)

	if e != nil {
		log.Fatalf("Failed to load configuration %s", e)
	}

	return &config
}
