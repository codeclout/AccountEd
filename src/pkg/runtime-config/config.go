package runtime_config

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

const configFileLocation string = "./runtime-config.hcl"

type AccessControlListConfig struct {
	Read  string `hcl:"read"`
	Write string `hcl:"write"`
}

type AWSRuntimeConfig struct {
	RoleDuration int32 `hcl:"role_duration"`
}

type AWSCloudConfig struct {
	AccessKeyId     *string
	AWSRegion       string
	AWSProfile      string
	Expiration      *time.Time
	RoleArn         string
	SecretAccessKey *string
	SessionName     string
	SessionToken    *string
}

type Cloud struct {
	AWS AWSRuntimeConfig `hcl:"aws,block"`
}

type RuntimeConfig struct {
	*AWSCloudConfig
	InfoLog *log.Logger

	CloudProvider Cloud `hcl:"cloudProvider,block"`
	Environment   string
	HostName      string

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
	UseCache                 bool   `hcl:"use_cache"`
}

func (r *RuntimeConfig) SetAwsSession(c sts.AssumeRoleOutput) error {
	*r.AWSCloudConfig.AccessKeyId = *c.Credentials.AccessKeyId
	*r.AWSCloudConfig.SecretAccessKey = *c.Credentials.SecretAccessKey
	*r.AWSCloudConfig.Expiration = *c.Credentials.Expiration
	*r.AWSCloudConfig.SessionToken = *c.Credentials.SessionToken

	return nil
}

func GetConfig() (*RuntimeConfig, error) {
	config := RuntimeConfig{
		Environment: os.Getenv("ENVIRONMENT"),
		AWSCloudConfig: &AWSCloudConfig{
			AWSRegion:   os.Getenv("AWS_REGION"),
			AWSProfile:  os.Getenv("AWS_PROFILE"),
			RoleArn:     os.Getenv("AWS_ROLE_ARN"),
			SessionName: os.Getenv("AWS_SESSION_NAME"),
		},
	}

	e := hclsimple.DecodeFile(configFileLocation, nil, &config)

	if e != nil {
		return nil, e
	}

	return &config, nil
}
