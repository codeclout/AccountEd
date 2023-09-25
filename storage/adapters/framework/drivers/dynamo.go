package drivers

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/api"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

var defaultRouteDuration = sessionTypes.DefaultRouteDuration(2000)

type cc = context.Context

type ConfirmedRegReq = pb.StoreConfirmedRegistrationRequest
type ConfirmedRegResp = pb.StoreConfirmedRegistrationResponse

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

func (a *Adapter) processFetchToken(in *pb.FetchTokenRequest) (*storageTypes.FetchTokenIn, error) {
	var creds credentials.StaticCredentialsProvider

	e := json.Unmarshal(in.GetCredentials(), &creds)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, status.Error(codes.Internal, e.Error())
	}

	return &storageTypes.FetchTokenIn{
		Credentials: creds,
		TableName:   in.GetTableName(),
		Token:       in.GetToken(),
	}, nil
}

func (a *Adapter) FetchToken(ctx cc, in *pb.FetchTokenRequest) (*pb.FetchTokenResponse, error) {
	apiData, e := a.processFetchToken(in)
	if e != nil {
		return nil, e
	}

	ch := make(chan *pb.FetchTokenResponse, 1)
	ech := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, in.GetToken())
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	a.dynamoApi.GetToken(ctx, apiData, ch, ech)

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

func (a *Adapter) processStorePublicToken(in *pb.TokenStoreRequest) (*storageTypes.TokenStorePayload, error) {
	var s string
	var creds credentials.StaticCredentialsProvider

	e := json.Unmarshal(in.GetSessionServiceAWScredentials(), &creds)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, status.Error(codes.Internal, e.Error())
	}

	if in.GetMemberId() == (s) ||
		in.GetPublicKey() == (s) ||
		in.GetSessionTableName() == (s) ||
		in.GetToken() == (s) ||
		in.GetTokenId() == (s) {
		return nil, status.Error(codes.InvalidArgument, "invalid token store request")
	}

	region, ok := a.config["AWSRegion"].(string)
	if !ok {
		a.monitor.LogGenericError("region not set in environment")
		return nil, status.Error(codes.Internal, "region not configured in environment")
	}

	return &storageTypes.TokenStorePayload{
		AWSRegion:        aws.String(region),
		Credentials:      creds,
		HasAutoCorrect:   in.GetHasAutoCorrect(),
		MemberId:         in.GetMemberId(),
		PublicKey:        in.GetPublicKey(),
		SessionTableName: in.GetSessionTableName(),
		Token:            in.GetToken(),
		TokenId:          in.GetTokenId(),
		Ttl:              in.GetTtl(),
	}, nil
}

func (a *Adapter) StorePublicToken(ctx cc, in *pb.TokenStoreRequest) (*pb.TokenStoreResponse, error) {
	apiData, e := a.processStorePublicToken(in)
	if e != nil {
		return nil, e
	}

	ch := make(chan *pb.TokenStoreResponse, 1)
	ech := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, in.GetTokenId())
	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()

	a.dynamoApi.CreatePublicTokenItem(ctx, apiData, ch, ech)

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

func (a *Adapter) StoreConfirmedRegistration(ctx cc, in *ConfirmedRegReq) (*ConfirmedRegResp, error) {
	return nil, errors.New("not implemented")
}
