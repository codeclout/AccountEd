package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/core"
	"github.com/codeclout/AccountEd/storage/ports/framework/driven"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type configMap = map[string]interface{}
type gClients = protocol.AdapterServiceClients
type icore = core.DynamoDbCorePort
type idriven = driven.DynamoDbDrivenPort
type preconfirmin = storageTypes.PreRegistrationSessionAPIin
type preconfirmResponsepb = pb.PreRegistrationConfirmationResponse

type Adapter struct {
	config     map[string]interface{}
	core       core.DynamoDbCorePort
	driven     driven.DynamoDbDrivenPort
	gRPClients *protocol.AdapterServiceClients
	monitor    monitoring.Adapter
}

func NewAdapter(config configMap, core icore, gcs *gClients, driven idriven, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:     config,
		core:       core,
		driven:     driven,
		gRPClients: gcs,
		monitor:    monitor,
	}
}

func (a *Adapter) PreRegistrationConfirmationApi(ctx context.Context, in preconfirmin, ch chan *preconfirmResponsepb, ech chan error) {
	region, ok := a.config["Region"].(string)
	if !ok {
		a.monitor.LogGenericError("region not set in environment")
		ech <- status.Error(codes.Internal, "region not configured in environment")
	}

	client, e := a.driven.GetDynamoClient(ctx, in.SessionServiceCredentials, &region)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		ech <- status.Error(codes.Internal, "unable to retrieve storage client")
	}

	result, e := a.driven.StoreSession(ctx, client, in)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		ech <- status.Error(codes.Internal, "unable to store session")
	}

	ctx = context.WithValue(ctx, "api_input", in)
	ctx = context.WithValue(ctx, "driven_output", result)

	out, e := a.core.PreRegistrationConfirmationCore(ctx)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		ech <- status.Error(codes.Internal, "internal error")
	}
	
	ch <- out
}
