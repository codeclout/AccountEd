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
