package member

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) GetAwsRegion() (string, error) {
	awsRegion, ok := a.config["Region"].(string)
	if !ok {
		return "", errors.New("Check the 'Region' in application configuration")
	}
	return awsRegion, nil
}

func (a *Adapter) createSecretsManagerClient(awsconfig []byte, awsRegion string) (*secretsmanager.Client, error) {
	var creds credentials.StaticCredentialsProvider

	err := json.Unmarshal(awsconfig, &creds)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling awsconfig: %w", err)
	}

	secretsClient := secretsmanager.NewFromConfig(aws.Config{Credentials: creds}, func(options *secretsmanager.Options) {
		options.Region = awsRegion
	})

	return secretsClient, nil
}

func (a *Adapter) createSystemsManagerClient(awsconfig []byte, awsRegion string) (*ssm.Client, error) {
	var creds credentials.StaticCredentialsProvider

	err := json.Unmarshal(awsconfig, &creds)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling awsconfig: %w", err)
	}

	ssmClient := ssm.NewFromConfig(aws.Config{Credentials: creds}, func(options *ssm.Options) {
		options.Region = awsRegion
	})

	return ssmClient, nil
}

// GetSessionIdKey retrieves a session ID key from AWS Secret Manager via the System Manager's Parameter Store.
// AWS region and 'PreRegistrationParameter' are extracted from the application configuration.
func (a *Adapter) GetSessionIdKey(ctx context.Context, awsconfig []byte) (*string, error) {
	var creds credentials.StaticCredentialsProvider

	e := json.Unmarshal(awsconfig, &creds)
	if e != nil {
		return nil, e
	}

	awsRegion, ok := a.config["Region"].(string)
	if !ok {
		return nil, errors.New("driven member -> AWS region missing | Check the 'Region' in application configuration")
	}

	ssmClient := ssm.NewFromConfig(aws.Config{Credentials: creds}, func(options *ssm.Options) {
		options.Region = awsRegion
	})
	secretsClient := secretsmanager.NewFromConfig(aws.Config{Credentials: creds}, func(options *secretsmanager.Options) {
		options.Region = awsRegion
	})

	name, ok := a.config["PreRegistrationParameter"].(string)
	if !ok {
		return nil, errors.New("driven member -> missing parameter in environment")
	}

	in := ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}

	parameter, e := ssmClient.GetParameter(ctx, &in)
	if e != nil {
		return nil, e
	}

	parameterValue := parameter.Parameter.Value
	if parameterValue == nil {
		return nil, errors.New("expected param value, got nil")
	}

	secIn := secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(*parameterValue),
		VersionStage: aws.String("AWSCURRENT"),
	}

	sec, e := secretsClient.GetSecretValue(ctx, &secIn)
	if e != nil {
		return nil, e
	}

	return sec.SecretString, nil
}
