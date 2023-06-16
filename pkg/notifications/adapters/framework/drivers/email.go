package drivers

import (
	"context"
	"errors"
	notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
	"strconv"
	"time"

	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/api"
)

type TransactionID string

type Adapter struct {
	apiEmail api.EmailApiPort
	config   map[string]interface{}
	log      *slog.Logger
}

func NewAdapter(api api.EmailApiPort, config map[string]interface{}, log *slog.Logger) *Adapter {
	return &Adapter{
		apiEmail: api,
		config:   config,
		log:      log,
	}
}

// ValidateEmailAddress takes a context and a ValidateEmailAddressRequest, sends the request using the EmailApiPort, and returns a
// ValidateEmailAddressResponse and an error if any. It sets a timeout for the request using the "sla_route_performance" config value, and listens for
// a response using channels. If the context times out, it returns a "request timeout" error. If an error is received from the error channel, it is
// logged and returned as is. Otherwise, the received response is returned with a success log message.
func (a *Adapter) ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error) {
	address := email.GetAddress()
	sla, ok := a.config["sla_route_performance"].(string)
	if !ok {
		a.log.Error("drivers -> static config sla_route_performance is not a string")
		return nil, notifications.ErrorStaticConfig(errors.New("wrong type: sla"))
	}

	b, _ := strconv.Atoi(sla)
	transactionID := TransactionID("transactionID")

	ch := make(chan *pb.ValidateEmailAddressResponse, 1)
	ctx = context.WithValue(ctx, transactionID, address)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(b)*time.Millisecond)

	defer cancel()

	errorch := make(chan error, 1)
	a.apiEmail.ValidateEmailAddress(ctx, address, ch, errorch)

	select {
	case <-ctx.Done():
		t := ctx.Value(transactionID)
		a.log.Error("request timeout", "transaction_id", t.(string))
		return nil, errors.New("request timeout")

	case out := <-ch:
		t := ctx.Value(transactionID)
		a.log.Info("success", "transaction_id", t.(string))
		return out, nil

	case e := <-errorch:
		t := ctx.Value(transactionID)
		a.log.Error(e.Error(), "transaction_id", t.(string))
		return nil, e
	}

}
