package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// FindParameter retrieves an AWS Systems Manager string parameter
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region
//	api is the interface that defines the method call
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a GetParameterOutput object containing the result of the service call and nil
//	Otherwise, nil and an error from the call to GetParameter
func (a *Adapter) FindParameter(c context.Context, api *ssm.Client, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	return api.GetParameter(c, input)
}

func (a *Adapter) GetParam(in *aws.Config, name *string) (*[]byte, error) {

	client := ssm.NewFromConfig(*in)
	out := &ssm.GetParameterInput{
		Name: name,
	}
	result, e := a.FindParameter(context.TODO(), client, out)

	if e != nil {
		a.log("fatal", e.Error())
	}

	b, _ := json.Marshal(result)

	return &b, nil
}

func (a *Adapter) GetSecret(in *aws.Config, id *string) (*string, error) {
	var (
		e        error
		secretId *ssm.GetParameterOutput
	)

	client := secretsmanager.NewFromConfig(*in)

	s, e := a.GetParam(in, id)
	if e != nil {
		a.log("fatal", e.Error())
	}

	e = json.Unmarshal(*s, &secretId)

	if e != nil {
		a.log("fatal", e.Error())
	}

	result, e := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(*secretId.Parameter.Value),
		VersionStage: aws.String("AWSCURRENT"),
	})

	if e != nil {
		a.log("fatal", e.Error())
	}

	return result.SecretString, nil
}
