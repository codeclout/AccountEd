package in

import "github.com/aws/aws-sdk-go-v2/aws"

type CredentialsPort interface {
	ExportCreds() *aws.Config
	ReLoadCreds() *aws.Config
}
