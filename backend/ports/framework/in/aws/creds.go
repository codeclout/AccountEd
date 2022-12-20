package aws

import "github.com/aws/aws-sdk-go-v2/aws"

type CredentialsPort interface {
	LoadCreds() *aws.Config
}
