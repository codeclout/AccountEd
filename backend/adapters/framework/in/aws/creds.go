package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type (
	c map[string]interface{}
	l func(level, msg string)
)

type Adapter struct {
	cloudConfig *aws.Config
	config      c
	log         l
}

func getCreds(logger l, roleArn string) *aws.Config {
	cfg, e := config.LoadDefaultConfig(context.TODO())
	if e != nil {
		logger("fatal", e.Error())
	}

	client := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(client, roleArn)
	cfg.Credentials = aws.NewCredentialsCache(creds)

	return &cfg
}

func NewAdapter(logger l, runtimeConfig c) *Adapter {
	var (
		roleArn = runtimeConfig["AwsRoleToAssume"].(string)
	)

	cfg := getCreds(logger, roleArn)

	return &Adapter{
		cloudConfig: cfg,
		config:      runtimeConfig,
		log:         logger,
	}
}

func (a *Adapter) LoadCreds() *aws.Config {
	return getCreds(a.log, a.config["AwsRoleToAssume"].(string))
}
