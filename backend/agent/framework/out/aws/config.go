package aws

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

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
	Organization string                  `hcl:"organization"`
}

func (r *RuntimeConfig) SetAwsSession(c sts.AssumeRoleOutput) error {
	*r.AWSCloudConfig.AccessKeyId = *c.Credentials.AccessKeyId
	*r.AWSCloudConfig.SecretAccessKey = *c.Credentials.SecretAccessKey
	*r.AWSCloudConfig.Expiration = *c.Credentials.Expiration
	*r.AWSCloudConfig.SessionToken = *c.Credentials.SessionToken

	return nil
}

func GetConfig(ctx context.Context) (*RuntimeConfig, error) {
	wd, wdError := os.Getwd()
	if wdError != nil {
		return nil, wdError
	}

	configFileLocation := filepath.Join(wd, "config.hcl")

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
