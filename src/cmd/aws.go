package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	runtimeconfig "github.com/codeclout/AccountEd/src/pkg/runtime-config"
)

// STSAssumeRoleAPI defines the interface for the AssumeRole function.
// We use this interface to test the function using a mocked service.
type STSAssumeRoleAPI interface {
	AssumeRole(ctx context.Context,
		params *sts.AssumeRoleInput,
		optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

// TakeRole gets temporary security credentials to access resources.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If successful, an AssumeRoleOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to AssumeRole.
func TakeRole(c context.Context, api STSAssumeRoleAPI, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return api.AssumeRole(c, input)
}

func AWSGetSessionToken() (*sts.AssumeRoleOutput, error) {
	rtc, e := runtimeconfig.GetConfig()
	if e != nil {
		return nil, e
	}

	if rtc.RoleArn == "" || rtc.SessionName == "" {
		log.Fatalf("You must supply a role ARN and session name")
	}

	awscfg, e := config.LoadDefaultConfig(context.TODO())
	if e != nil {
		panic("configuration error, " + e.Error())
	}

	client := sts.NewFromConfig(awscfg)
	input := &sts.AssumeRoleInput{
		RoleArn:         &rtc.RoleArn,
		RoleSessionName: &rtc.SessionName,
	}

	result, ests := TakeRole(context.TODO(), client, input)
	if ests != nil {
		log.Fatalf(ests.Error())
	}

	return result, nil
}
