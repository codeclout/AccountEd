package aws

import "github.com/aws/aws-sdk-go-v2/aws"

type ParameterPort interface {
	GetParam(config *aws.Config, name *string) (*[]byte, error)
	GetSecret(config *aws.Config, id *string) (*string, error)
}
