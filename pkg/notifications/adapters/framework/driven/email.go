package driven

import (
	"context"
	"encoding/json"
	"golang.org/x/exp/slog"
	"net/http"

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

// EmailVerificationProcessor is a method of the Adapter struct which takes a context and an EmailDrivenIn object as input and returns
// a ValidateEmailOut object and an error. It sends an API request to verify the email address by making an HTTP GET request with the appropriate
// email_processor_api_key and email address. If the response is successful, it decodes the JSON response body into a ValidateEmailOut object, otherwise
// it returns an error. The error can be caused by invalid input data or an unsuccessful response with a status code higher than 299.
func (a *Adapter) EmailVerificationProcessor(ctx context.Context, in *notifications.EmailDrivenIn) (*notifications.ValidateEmailOut, error) {
	var out notifications.ValidateEmailOut

	client := &http.Client{}

	req, _ := http.NewRequest("GET", in.Endpoint, nil)
	params := req.URL.Query()

	emailProcessorApiKey, ok := a.config["email_processor_api_key"].(string)
	if !ok {
		a.log.Error("driven -> email processor api emailProcessorApiKey is not a string")
		return nil, notifications.ErrorStaticConfig(errors.New("core -> wrong type: emailProcessorApiKeyx"))
	}

	params.Add("api_key", emailProcessorApiKey)
	params.Add("email", in.EmailAddress)

	req.URL.RawQuery = params.Encode()

	response, e := client.Do(req)
	if e != nil || response.StatusCode > 299 {
		return nil, errors.New(response.Status)
	}

	e = json.NewDecoder(response.Body).Decode(&out)
	return &out, nil
}
