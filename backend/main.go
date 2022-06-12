package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func createEnvironmentSession() *sts.AssumeRoleOutput {
	awc, e := AWSGetSessionToken()
	if e != nil {
		log.Fatalf("Unable to retrieve environment session - failed with error: %s", e)
	}

	return awc
}

func main() {
	ctx := context.Background()
	// app := fiber.New(fiber.Config{})
	//
	//fx.New(
	//	fx.Provide(app),
	//).Run()

	c, e := aws.GetConfig()

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
