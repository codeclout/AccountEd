package driven

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type ValidateEmailIn = notifications.ValidateEmailIn
type ValidateEmailOut = notifications.ValidateEmailOut

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(c map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  c,
		monitor: monitor,
	}
}

func (a *Adapter) EmailVerificationProcessor(ctx context.Context, in *ValidateEmailIn) (*ValidateEmailOut, error) {
	var out notifications.ValidateEmailOut

	endpoint := fmt.Sprintf("%s%s", in.ProcessorDomain, in.ProcessorEndpoint)
	client := &http.Client{}

	req, e := http.NewRequest("GET", endpoint, nil)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}

	params := req.URL.Query()

	params.Add("api_key", in.ProcessorKey)
	params.Add("email", in.Address)
	req.URL.RawQuery = params.Encode()

	response, e := client.Do(req)
	if e != nil || response.StatusCode > 299 {
		a.monitor.LogGrpcError(ctx, fmt.Sprintf("email processor api returned -> %s", response.Status))
		return nil, errors.New(response.Status)
	}

	defer func(Body io.ReadCloser) {
		e := Body.Close()
		if e != nil {
			a.monitor.LogGenericError(e.Error())
		}
	}(response.Body)

	e = json.NewDecoder(response.Body).Decode(&out)
	if e != nil {
		a.monitor.LogGrpcError(ctx, "driven EmailVerificationProcessor -> unable to decode response body")
		return nil, errors.New("unable to decode response body")
	}

	return &out, nil
}
