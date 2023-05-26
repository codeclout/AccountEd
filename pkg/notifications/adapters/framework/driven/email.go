package driven

import (
  "bytes"
  "context"
  "encoding/json"
  "mime/multipart"
  "net/http"

  notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
)

type Adapter struct {
  config map[string]interface{}
}

func NewAdapter(c map[string]interface{}) *Adapter {
  return &Adapter{config: c}
}

func (a *Adapter) EmailVerificationProcessor(ctx context.Context, reqBody *bytes.Buffer, writer *multipart.Writer) (*notifications.ValidateEmailOut, error) {
  var out notifications.ValidateEmailOut
  client := &http.Client{}

  req, _ := http.NewRequest("POST", a.config["email_processor_domain"].(string)+"/v4/address/validate", reqBody)
  req.Header.Set("Content-Type", writer.FormDataContentType())
  req.SetBasicAuth("api", a.config["email_processor_api_key"].(string))

  response, e := client.Do(req)
  if e != nil {
    return nil, e
  }

  e = json.NewDecoder(response.Body).Decode(&out)
  return &out, nil
}
