package drivers

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/api"
)

var defaultRouteDuration = sessiontypes.DefaultRouteDuration(2000)

type Adapter struct {
	AwsGrpcClient awspb.AWSResourceClientServiceClient
	config        map[string]interface{}
	dynamoApi     api.DynamoDbApiPort
	monitor       monitoring.Adapter
	waitgroup     *sync.WaitGroup
}

func NewAdapter(config map[string]interface{}, dynamoApi api.DynamoDbApiPort, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	return &Adapter{
		config:    config,
		dynamoApi: dynamoApi,
		monitor:   monitor,
		waitgroup: wg,
	}
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
		return 0, errors.New("error converting slaroutperformance to int: " + e.Error())
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

func (a *Adapter) GetPreRegistrationBySessionId(ctx context.Context, request *pb.PreRegistrationConfirmationRequest) (*pb.PreRegistrationConfirmationResponse, error) {
	sessionID := request.GetSessionID()

	ch := make(chan *pb.PreRegistrationConfirmationResponse, 1)
	errorch := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, sessionID)
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	a.dynamoApi.PreRegistrationConfirmationApi(ctx, sessionID, ch, errorch)

	select {
	case <-ctx.Done():
		a.monitor.LogGrpcError(ctx, "request timeout")
		return nil, errors.New("request timeout")

	case out := <-ch:
		t := ctx.Value(a.monitor.LogLabelTransactionID)
		a.monitor.LogGrpcInfo(ctx, fmt.Sprintf("pre registration confirmation success for %s", t))
		return out, nil

	case e := <-errorch:
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}
}
