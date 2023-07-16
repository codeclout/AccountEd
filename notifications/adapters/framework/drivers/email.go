package drivers

import (
	"context"
	"errors"
	"strconv"
	"time"

	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/notifications/gen/v1"
	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	"github.com/codeclout/AccountEd/notifications/ports/api"
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

func (a *Adapter) setContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	deadline, ok := ctx.Deadline()

	if ok {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		return ctx, cancel
	}

	sla, e := a.getRequestSLA()
	if e != nil {
		sla = int(defaultRouteDuration)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(sla)*time.Millisecond)
	return ctx, cancel
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error) {
	address := email.GetAddress()

	ch := make(chan *pb.ValidateEmailAddressResponse, 1)
	errorch := make(chan error, 1)

	ctx = context.WithValue(ctx, transactionID, address)
	ctx, cancel := a.setContextTimeout(ctx)

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

func (a *Adapter) SendPreRegistrationEmail(ctx context.Context, in *pb.NoReplyEmailNotificationRequest) (*pb.NoReplyEmailNotificationResponse, error) {
	awscredentials := in.GetAwsCredentials()
	domain := in.GetDomain()
	fromAddress := in.GetFromAddress()
	sessionID := in.GetSessionId()
	toAddress := in.GetToAddress()

	ch := make(chan *pb.NoReplyEmailNotificationResponse, 1)
	errorch := make(chan error, 1)

	ctx = context.WithValue(ctx, transactionID, sessionID)
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	apiData := notifications.NoReplyEmailIn{
		AWSCredentials: awscredentials,
		Domain:         domain,
		FromAddress:    fromAddress,
		SessionID:      sessionID,
		ToAddress:      toAddress,
	}

	a.apiEmail.SendPreRegistrationEmailAPI(ctx, &apiData, ch, errorch)

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
