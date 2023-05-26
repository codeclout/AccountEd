package drivers

import (
	"context"
	"errors"
	"strconv"
	"time"

	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/api"
)

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

func (a *Adapter) ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error) {
	address := email.GetAddress()
	b, _ := strconv.Atoi(a.config["sla_routePerformance"].(string))

	ch := make(chan *pb.ValidateEmailAddressResponse, 1)
	ctx = context.WithValue(ctx, "transactionID", address)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(b)*time.Millisecond)

	defer cancel()

	errorch := make(chan error, 1)
	a.apiEmail.ValidateEmailAddress(ctx, address, ch, errorch)

	select {
	case <-ctx.Done():
		a.log.ErrorCtx(ctx, "request timeout")
		return nil, errors.New("request timeout")

	case out := <-ch:
		a.log.InfoCtx(ctx, "success")
		return out, nil

	case e := <-errorch:
		a.log.ErrorCtx(ctx, "error")
		return nil, e
	}

}
