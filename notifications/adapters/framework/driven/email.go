package driven

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"golang.org/x/exp/slog"

	"github.com/pkg/errors"

	notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type Adapter struct {
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(c map[string]interface{}, log *slog.Logger) *Adapter {
	return &Adapter{
		config: c,
		log:    log,
	}
}

func (a *Adapter) getPreRegistrationNoReplyContent(body, subject string) *types.EmailContent {
	out := types.EmailContent{
		Simple: &types.Message{
			Body: &types.Body{
				Text: &types.Content{
					Data:    aws.String(body),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &types.Content{
				Data:    aws.String(subject),
				Charset: aws.String("UTF-8"),
			},
		},
	}

	return &out
}

func (a *Adapter) EmailVerificationProcessor(ctx context.Context, in *notifications.EmailDrivenIn) (*notifications.ValidateEmailOut, error) {
	var out notifications.ValidateEmailOut

	client := &http.Client{}

	req, e := http.NewRequest("GET", in.Endpoint, nil)
	if e != nil {
		a.log.Error(e.Error())
		return nil, e
	}

	params := req.URL.Query()

	emailProcessorApiKey, ok := a.config["EmailProcessorAPIKey"].(string)
	if !ok {
		a.log.Error("driven -> email processor api emailProcessorApiKey is not a string")
		return nil, notifications.ErrorStaticConfig(errors.New("core -> wrong type: emailProcessorApiKey"))
	}

	params.Add("api_key", emailProcessorApiKey)
	params.Add("email", in.EmailAddress)
	req.URL.RawQuery = params.Encode()

	response, e := client.Do(req)
	if e != nil || response.StatusCode > 299 {
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New(response.Status))
	}

	e = json.NewDecoder(response.Body).Decode(&out)
	if e != nil {
		a.log.Error("driven -> unable to decode response body")
		return nil, notifications.ErrorEmailVerificationProcessor(errors.New("unable to decode response body"))
	}

	return &out, nil
}

func (a *Adapter) SendPreRegistrationEmail(ctx context.Context, awsconfig []byte, body, subject string, in *notifications.NoReplyEmailIn) (*notifications.NoReplyEmailOut, error) {
	var creds credentials.StaticCredentialsProvider
	var out notifications.NoReplyEmailOut

	e := json.Unmarshal(awsconfig, &creds)
	if e != nil {
		return nil, e
	}

	awsRegion, ok := a.config["Region"].(string)
	if !ok {
		return nil, errors.New("driven member -> AWS region missing | Check the 'Region' in application configuration")
	}

	emailClient := sesv2.NewFromConfig(aws.Config{Credentials: creds}, func(options *sesv2.Options) {
		options.Region = awsRegion
	})

	x, e := emailClient.SendEmail(ctx, &sesv2.SendEmailInput{
		Content: a.getPreRegistrationNoReplyContent(body, subject),
		Destination: &types.Destination{
			ToAddresses: in.ToAddress,
		},
		FromEmailAddress: aws.String(in.FromAddress),
	})

	if e != nil {
		a.log.Error(e.Error())
		return nil, e
	}

	out = notifications.NoReplyEmailOut{MessageID: *x.MessageId}
	return &out, nil
}
