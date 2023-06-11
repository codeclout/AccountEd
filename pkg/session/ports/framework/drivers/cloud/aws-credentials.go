package cloud

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type CredentialsAWSPort interface {
	AssumeRoleCredentials(ctx context.Context, arn string) (*aws.Config, error)
}
