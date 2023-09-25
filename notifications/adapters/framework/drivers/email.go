package drivers

import (
	"context"
	"errors"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	"github.com/codeclout/AccountEd/notifications/ports/api"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

var defaultRouteDuration = notifications.DefaultRouteDuration(2000)

type Adapter struct {
	apiEmail api.EmailApiPort
	config   map[string]interface{}
	monitor  monitoring.Adapter
}

func NewAdapter(api api.EmailApiPort, config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		apiEmail: api,
		config:   config,
		monitor:  monitor,
	}
}

func (a *Adapter) getConfigString(key, msg string) (string, error) {
	item, ok := a.config[key].(string)
	if !ok {
		a.monitor.LogGenericError(msg)
		return "", errors.New(msg)
	}

	return item, nil
}

func (a *Adapter) getRequestSLA() (int, error) {
	sla, ok := a.config["SLARoutePerformance"].(string)
	if !ok {
		a.monitor.LogGenericError("drivers -> static config sla_route_performance is not a string")
		return 0, errors.New("wrong type: sla")
	}

	i, e := strconv.Atoi(sla)
	if e != nil {
		a.monitor.LogGenericError("drivers -> error converting sla_route_performance to int")
		return 0, errors.New("error converting SLARoutePerformance to int: " + e.Error())
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

func (a *Adapter) processValidateEmail(in *pb.ValidateEmailAddressRequest) (*notifications.ValidateEmailIn, error) {
	var s string

	if in.GetAddress() == (s) {
		return nil, status.Error(codes.InvalidArgument, "invalid email address")
	}

	processorDomain, e := a.getConfigString("EmailProcessorDomain", "email verification domain is invalid")
	if e != nil {
		return nil, e
	}

	processorEndpoint, e := a.getConfigString("EmailVerifierApiPath", "email verification processor endpoint is invalid")
	if e != nil {
		return nil, e
	}

	processorKey, e := a.getConfigString("EmailProcessorAPIKey", "EmailProcessorAPIKey is not a string")
	if e != nil {
		return nil, e
	}

	return &notifications.ValidateEmailIn{
		Address:           in.GetAddress(),
		ProcessorDomain:   processorDomain,
		ProcessorEndpoint: processorEndpoint,
		ProcessorKey:      processorKey,
	}, nil
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error) {
	metadata, e := a.processValidateEmail(email)
	if e != nil {
		return nil, status.Error(codes.FailedPrecondition, e.Error())
	}

	ch := make(chan *pb.ValidateEmailAddressResponse, 1)
	ech := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, metadata.Address)
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	a.apiEmail.ValidateEmailAddress(ctx, metadata, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogGrpcError(ctx, "request timeout")
		return nil, errors.New("request timeout")

	case out := <-ch:
		a.monitor.LogGrpcInfo(ctx, "success")
		return out, nil

	case e := <-ech:
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}

}

func (a *Adapter) processRegistrationEmail(in *pb.NoReplyEmailNotificationRequest) (*notifications.NoReplyEmailIn, error) {
	var s string

	if in.Domain == (s) || in.FromAddress == (s) || in.Token == (s) {
		return nil, status.Error(codes.InvalidArgument, "invalid registration email request")
	}

	return &notifications.NoReplyEmailIn{
		AWSCredentials: in.GetAwsCredentials(),
		Domain:         in.GetDomain(),
		FromAddress:    in.GetFromAddress(),
		Token:          in.GetToken(),
		ToAddress:      in.GetToAddress(),
	}, nil
}

func (a *Adapter) SendPreRegistrationEmail(ctx context.Context, in *pb.NoReplyEmailNotificationRequest) (*pb.NoReplyEmailNotificationResponse, error) {
	metadata, e := a.processRegistrationEmail(in)
	if e != nil {
		return nil, status.Error(codes.FailedPrecondition, e.Error())
	}

	ch := make(chan *pb.NoReplyEmailNotificationResponse, 1)
	ech := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, metadata.Token)
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	a.apiEmail.SendPreRegistrationEmailAPI(ctx, metadata, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogGrpcError(ctx, "request timeout")
		return nil, errors.New("request timeout")

	case out := <-ch:
		a.monitor.LogGrpcInfo(ctx, "success")
		return out, nil

	case e := <-ech:
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}
}
