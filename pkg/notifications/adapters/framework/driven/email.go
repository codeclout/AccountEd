package driven

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type Adapter struct {
	config map[string]interface{}
}

func NewAdapter(c map[string]interface{}) *Adapter {
	return &Adapter{config: c}
}

// EmailVerificationProcessor takes a context and an EmailDrivenIn object as input, and returns a ValidateEmailOut object and an error if there
// any issues during processing. It sends an API request to verify the email address by making an HTTP GET request with the email_processor_api_key and
// email address. If the response is successful, it decodes the JSON response body into a ValidateEmailOut object and returns it along with any errors.
func (a *Adapter) EmailVerificationProcessor(ctx context.Context, in *notifications.EmailDrivenIn) (*notifications.ValidateEmailOut, error) {
	var out notifications.ValidateEmailOut

	client := &http.Client{}

	req, _ := http.NewRequest("GET", in.Endpoint, nil)
	params := req.URL.Query()

	params.Add("api_key", a.config["email_processor_api_key"].(string))
	params.Add("email", in.EmailAddress)

	req.URL.RawQuery = params.Encode()

	response, e := client.Do(req)
	if e != nil || response.StatusCode > 299 {
		return nil, errors.New(response.Status)
	}

	e = json.NewDecoder(response.Body).Decode(&out)
	return &out, nil
}
