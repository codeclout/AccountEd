package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type c map[string]interface{}
type l func(level, msg string)

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

// TakeRole gets temporary security credentials to access resources.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If successful, an AssumeRoleOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to AssumeRole.
func TakeRole(c context.Context, api *sts.Client, in *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return api.AssumeRole(c, in)
}

func (a *Adapter) LoadCreds() ([]byte, error) {
	var (
		roleArn     string = a.config["AwsRoleToAssume"].(string)
		sessionName string = "user-account-type-session"
	)

	cfg, e := config.LoadDefaultConfig(context.TODO())

	if e != nil {
		a.log("fatal", e.Error())
	}

	client := sts.NewFromConfig(cfg)
	in := &sts.AssumeRoleInput{
		RoleArn:         &roleArn,
		RoleSessionName: &sessionName,
	}

	result, e := TakeRole(context.TODO(), client, in)

	if e != nil {
		a.log("fatal", e.Error())
	}

	b, _ := json.Marshal(result)
	return b, nil
}
