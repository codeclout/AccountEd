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

type cc = context.Context
type configMap = map[string]interface{}
type gClients = protocol.AdapterServiceClients

type dynamoCore = core.DynamoDbCorePort
type drivenDynamo = driven.DynamoDbDrivenPort

type monitoringA = monitoring.Adapter

type TokenStorePayload = storageTypes.TokenStorePayload

type Adapter struct {
	config     configMap
	core       core.DynamoDbCorePort
	driven     driven.DynamoDbDrivenPort
	gRPClients *protocol.AdapterServiceClients
	monitor    monitoring.Adapter
}

func NewAdapter(config configMap, core dynamoCore, gcs *gClients, driven drivenDynamo, monitor monitoringA) *Adapter {
	return &Adapter{
		config:     config,
		core:       core,
		driven:     driven,
		gRPClients: gcs,
		monitor:    monitor,
	}
}

func (a *Adapter) CreatePublicTokenItem(ctx cc, in *TokenStorePayload, ch chan *pb.TokenStoreResponse, ech chan error) {
	client, e := a.driven.GetDynamoClient(ctx, in.Credentials, in.AWSRegion)
	if e != nil {
		ech <- status.Error(codes.Internal, "unable to retrieve storage client")
		return
	}

	result, e := a.driven.StoreToken(ctx, client, in)
	if e != nil {
		ech <- status.Error(codes.Internal, "unable to store session")
		return
	}

	out, e := a.core.ProcessStoredToken(ctx, result)
	if e != nil {
		ech <- status.Error(codes.Internal, "internal error")
		return
	}

	ch <- out
}

func (a *Adapter) GetToken(ctx cc, in *storageTypes.FetchTokenIn, ch chan *pb.FetchTokenResponse, ech chan error) {
	region, ok := a.config["AWSRegion"].(string)
	if !ok {
		a.monitor.LogGenericError("region not set in environment")
		ech <- status.Error(codes.Internal, "region not configured in environment")
		return
	}

	client, e := a.driven.GetDynamoClient(ctx, in.Credentials, &region)
	if e != nil {
		ech <- status.Error(codes.Internal, "unable to retrieve storage client")
		return
	}

	result, e := a.driven.GetTokenItem(ctx, client, *in)
	if e != nil {
		ech <- status.Error(codes.Internal, "unable to store session")
		return
	}

	out, e := a.core.ProcessFetchedToken(ctx, *result)
	if e != nil {
		ech <- status.Error(codes.Internal, "internal error")
		return
	}

	ch <- out

}
