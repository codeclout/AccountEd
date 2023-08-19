package drivers

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/api"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

var defaultRouteDuration = sessionTypes.DefaultRouteDuration(2000)

type cctx = context.Context

type ConfirmedRegReq = pb.StoreConfirmedRegistrationRequest
type ConfirmedRegResp = pb.StoreConfirmedRegistrationResponse
type PreRegReq = pb.PreRegistrationConfirmationRequest
type PreRegResp = pb.PreRegistrationConfirmationResponse

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

func (a *Adapter) getRequestSLA() (float64, error) {
	sla, ok := a.config["SLARoutes"].(float64)
	if !ok {
		a.monitor.LogGenericError("drivers -> static config sla_route_performance is not a string")
		return 0, errors.New("wrong type: sla")
	}

	return sla, nil
}

func (a *Adapter) setContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	deadline, ok := ctx.Deadline()

	if ok {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		return ctx, cancel
	}

	sla, e := a.getRequestSLA()
	if e != nil {
		sla = float64(defaultRouteDuration)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(sla)*time.Millisecond)
	return ctx, cancel
}

func (a *Adapter) StorePreConfirmationRegistrationSession(ctx cctx, request *PreRegReq) (*PreRegResp, error) {
	var apidata storageTypes.PreRegistrationSessionAPIin
	var credentialsOut credentials.StaticCredentialsProvider

	e := json.Unmarshal(request.GetSessionServiceAWScredentials(), &credentialsOut)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Internal, e.Error())
	}

	apidata = storageTypes.PreRegistrationSessionAPIin{
		AssociatedData:            request.GetAssociatedData(),
		EncryptedSessionID:        request.GetEncryptedSessionID(),
		ForwardedIP:               request.GetForwardedIp(),
		HasAutoCorrect:            request.GetHasAutoCorrect(),
		MemberID:                  request.GetMemberId(),
		Nonce:                     request.GetNonce(),
		SessionID:                 request.GetSessionID(),
		SessionServiceCredentials: &credentials.StaticCredentialsProvider{Value: credentialsOut.Value},
		SessionStorageTableName:   request.GetSessionTableName(),
		TTL:                       request.GetTtl(),
	}

	ch := make(chan *pb.PreRegistrationConfirmationResponse, 1)
	ech := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, request.GetSessionID())
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	a.dynamoApi.PreRegistrationConfirmationApi(ctx, apidata, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogGrpcError(ctx, "request timeout")
		return nil, status.Error(codes.DeadlineExceeded, "request timeout")

	case out := <-ch:
		t := ctx.Value(a.monitor.LogLabelTransactionID)
		a.monitor.LogGrpcInfo(ctx, fmt.Sprintf("pre registration confirmation success for %s", t))
		return out, nil

	case e := <-ech:
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Internal, e.Error())
	}
}

func (a *Adapter) StoreConfirmedRegistration(context.Context, *ConfirmedRegReq) (*ConfirmedRegResp, error) {
	return nil, errors.New("not implemented")
}
