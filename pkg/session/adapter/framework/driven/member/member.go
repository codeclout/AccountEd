package member

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

type Adapter struct {
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, memberDrivenLog *slog.Logger) *Adapter {
	return &Adapter{
		config: config,
		log:    memberDrivenLog,
	}
}

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
