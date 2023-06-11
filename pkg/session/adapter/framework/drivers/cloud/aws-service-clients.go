package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

type ErrorDefaultConfiguration = error

type Adapter struct {
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, log *slog.Logger) *Adapter {
	return &Adapter{
		config: config,
		log:    log,
	}
}

// AssumeRoleCredentials attempts to assume the specified AWS role and returns an AWS Config object with the
// assumed role's credentials, or an error if the operation failed. The function takes a context.Context and an IAM Role
// Amazon Resource Name (ARN) string as arguments, and returns a pointer to an aws.Config object and an error if any.
//
// The ARN parameter is the Amazon Resource Name for the role you want to assume. The context.Context is used for request
// cancellation and timeouts.
func (a *Adapter) AssumeRoleCredentials(ctx context.Context, arn, region *string) (*aws.Config, error) {
	configloader, e := config.LoadDefaultConfig(ctx, config.WithRegion(*region))
	if e != nil {
		return nil, ErrorDefaultConfiguration(errors.New("unable to load AWS configuration"))
	}

	client := sts.NewFromConfig(configloader)
	credentials := stscreds.NewAssumeRoleProvider(client, *arn)

	configloader.Credentials = aws.NewCredentialsCache(credentials)
	return &configloader, nil
}

// GetSystemsManagerClient creates and returns a new AWS Systems Manager (SSM) client instance using the provided AWS
// configuration. It takes a context.Context and a pointer to an aws.Config as arguments, and returns a
// pointer to an ssm.Client. The context.Context is used for request cancellation and timeouts, while the
// aws.Config should contain the necessary settings and credentials for connecting to the AWS API.
func (a *Adapter) GetSystemsManagerClient(ctx context.Context, config *aws.Config) *ssm.Client {
	client := ssm.NewFromConfig(*config)
	return client
}

// GetSecretsManagerClient creates and returns a new AWS Secrets Manager client instance using the provided AWS
// configuration. The function takes a context.Context and a pointer to an aws.Config as arguments and returns a
// pointer to a secretsmanager.Client. The context.Context is used for request cancellation and timeouts, while the
// aws.Config should contain the necessary settings and credentials for connecting to the AWS API.
func (a *Adapter) GetSecretsManagerClient(ctx context.Context, config *aws.Config) *secretsmanager.Client {
	client := secretsmanager.NewFromConfig(*config)
	return client
}
