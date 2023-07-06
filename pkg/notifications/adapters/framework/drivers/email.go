package drivers

import (
	"context"
	"errors"
	"strconv"
	"time"

	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	notifications "github.com/codeclout/AccountEd/pkg/notifications/notification-types"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/api"
)

var defaultRouteDuration = notifications.DefaultRouteDuration(2000)
var transactionID = notifications.TransactionID("transactionID")

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

func (a *Adapter) getRequestSLA() (int, error) {
	sla, ok := a.config["SLARoutePerformance"].(string)
	if !ok {
		a.log.Error("drivers -> static config sla_route_performance is not a string")
		return 0, notifications.ErrorStaticConfig(errors.New("wrong type: sla"))
	}

	i, e := strconv.Atoi(sla)
	if e != nil {
		a.log.Error("drivers -> error converting sla_route_performance to int")
		return 0, notifications.ErrorStaticConfig(errors.New("error converting slaroutperformance to int: " + e.Error()))
	}

	return i, nil
}

// ValidateEmailAddress takes a context and a ValidateEmailAddressRequest, sends the request using the EmailApiPort, and returns a
// ValidateEmailAddressResponse and an error if any. It sets a timeout for the request using the "sla_route_performance" config value, and listens for
// a response using channels. If the context times out, it returns a "request timeout" error. If an error is received from the error channel, it is
// logged and returned as is. Otherwise, the received response is returned with a success log message.
func (a *Adapter) ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error) {
	address := email.GetAddress()

	ch := make(chan *pb.ValidateEmailAddressResponse, 1)
	errorch := make(chan error, 1)

	ctx = context.WithValue(ctx, transactionID, address)
	sla, e := a.getRequestSLA()
	if e != nil {
		sla = int(defaultRouteDuration)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(sla)*time.Millisecond)

	defer cancel()

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
