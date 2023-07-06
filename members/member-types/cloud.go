package membertypes

import "github.com/aws/aws-sdk-go-v2/aws"

type CredentialsAWS struct {
	Value
}

type Value struct {
	aws.Credentials
}
