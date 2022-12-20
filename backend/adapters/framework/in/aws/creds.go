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
	config c
	log    l
}

func NewAdapter(logger l, runtimeConfig c) *Adapter {
	return &Adapter{
		config: runtimeConfig,
		log:    logger,
	}
}

func (a *Adapter) LoadCreds() *aws.Config {
	var (
		roleArn = a.config["AwsRoleToAssume"].(string)
	)

	cfg, e := config.LoadDefaultConfig(context.TODO())

	if e != nil {
		a.log("fatal", e.Error())
	}

	client := sts.NewFromConfig(cfg)

	creds := stscreds.NewAssumeRoleProvider(client, roleArn)
	cfg.Credentials = aws.NewCredentialsCache(creds)

	return &cfg
}
