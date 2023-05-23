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
	apiNotificationEmail api.EmailNotificationAPI
	config               map[string]string
	log                  *slog.Logger
}

func (a *Adapter) NewAdapter(api api.EmailNotificationAPI, config map[string]string, log *slog.Logger) *Adapter {
	return &Adapter{
		apiNotificationEmail: api,
		config:               config,
		log:                  log,
	}
}

func (a *Adapter) HandleValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error) {
	address := email.GetAddress()
	b, _ := strconv.Atoi(a.config["sla_routePerformance"])

	ch := make(chan *pb.ValidateEmailAddressResponse, 1)
	ctx = context.WithValue(ctx, "transactionID", address)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(b)*time.Millisecond)

	defer cancel()

	errorch := make(chan error, 1)
	a.apiNotificationEmail.ValidateEmailAddress(ctx, address, ch, errorch)

	select {
	case <-ctx.Done():
		a.log.ErrorCtx(ctx, "request timeout")
		return nil, errors.New("request timeout")

	case out := <-ch:
		a.log.InfoCtx(ctx, "success")
		return out, nil
	}
}
