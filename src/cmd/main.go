package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/codeclout/AccountEd/src/pkg/runtime-config"
)

func createEnvironmentSession() *sts.AssumeRoleOutput {
	awc, e := AWSGetSessionToken()
	if e != nil {
		log.Fatalf("Unable to retrieve evironment session - failed with error: %s", e)
	}

	return awc
}

func main() {
	c, e := runtime_config.GetConfig()

	if e != nil {
		log.Fatalf("Failed to load configuration %s", e)
	}

	cloudSession := createEnvironmentSession()

	c.AccessKeyId = cloudSession.Credentials.AccessKeyId
	c.Expiration = cloudSession.Credentials.Expiration
	c.SecretAccessKey = cloudSession.Credentials.SecretAccessKey
	c.SessionToken = cloudSession.Credentials.SessionToken

	fmt.Printf("Configuration is %#v", *c.AWSCloudConfig.SessionToken)
}
