package protocol

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	emailv1 "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers"
	awsv1 "github.com/codeclout/AccountEd/session/gen/aws/v1"
	membersessionv1 "github.com/codeclout/AccountEd/session/gen/members/v1"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

var ctx = context.Background()

type cfg = map[string]interface{}

type AdapterServiceClients struct {
	AwsSessionclient        *awsv1.AWSResourceClientServiceClient
	EmailNotificationclient *emailv1.EmailNotificationServiceClient
	MemberSessionclient     *membersessionv1.MemberSessionClient
	SessionStorageclient    *dynamov1.DynamoDBStorageServiceClient
	config                  cfg
	clientList              []*grpc.ClientConn
	internalConfiguration   *cfg
	monitor                 monitoring.Adapter
	wg                      *sync.WaitGroup
}

func NewGrpcAdapter(config cfg, monitor monitoring.Adapter, wg *sync.WaitGroup) *AdapterServiceClients {
	aServer := drivers.NewAdapter(monitor)
	internals := aServer.LoadServerConfiguration()

	return &AdapterServiceClients{
		config:                config,
		internalConfiguration: internals,
		monitor:               monitor,
		wg:                    wg,
	}
}

func (a *AdapterServiceClients) gRPCPostInit() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	a.monitor.Logger.Warn("closing grpc client connections")
}

func (a *AdapterServiceClients) initializeNotificationsClient() {
	var options []grpc.DialOption

	pConfig := *a.internalConfiguration
	to, ok := pConfig["GRPCClientConnectionTimeout"].(float64)

	if !ok {
		a.monitor.LogGenericError("sla_routes not configured")
		to = float64(2000)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(to)*time.Millisecond)
	defer cancel()

	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connection, e := grpc.DialContext(ctx, ":9000", options...)
	if e != nil {
		a.monitor.LogGenericError("gRPC notifications client failed to connect")
		os.Exit(1)
	}

	_ = append(a.clientList, connection)
	emailNotificationsClient := emailv1.NewEmailNotificationServiceClient(connection)

	a.EmailNotificationclient = &emailNotificationsClient
}

func (a *AdapterServiceClients) initializeSessionClient() {
	var options []grpc.DialOption

	pConfig := *a.internalConfiguration
	to, ok := pConfig["GRPCClientConnectionTimeout"].(float64)

	if !ok {
		a.monitor.LogGenericError("sla_routes not configured")
		to = float64(2000)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(to)*time.Millisecond)
	defer cancel()

	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connection, e := grpc.DialContext(ctx, ":9001", options...)
	if e != nil {
		a.monitor.LogGenericError("gRPC session client failed to connect")
		os.Exit(1)
	}

	_ = append(a.clientList, connection)
	awsSessionClient := awsv1.NewAWSResourceClientServiceClient(connection)
	memberSessionClient := membersessionv1.NewMemberSessionClient(connection)

	a.AwsSessionclient = &awsSessionClient
	a.MemberSessionclient = &memberSessionClient
}

func (a *AdapterServiceClients) initializeStorageClient() {
	var options []grpc.DialOption

	pConfig := *a.internalConfiguration
	to, ok := pConfig["GRPCClientConnectionTimeout"].(float64)

	if !ok {
		a.monitor.LogGenericError("sla_routes not configured")
		to = float64(2000)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(to)*time.Millisecond)
	defer cancel()

	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connection, e := grpc.DialContext(ctx, ":9003", options...)
	if e != nil {
		a.monitor.LogGenericError("gRPC session client failed to connect")
		os.Exit(1)
	}

	_ = append(a.clientList, connection)
	sessionStorageClient := dynamov1.NewDynamoDBStorageServiceClient(connection)

	a.SessionStorageclient = &sessionStorageClient
}

func (a *AdapterServiceClients) CloseClientConnection(conn *grpc.ClientConn) {
	e := conn.Close()
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}
}

func (a *AdapterServiceClients) InitializeClientsForMember() {
	defer a.wg.Done()

	a.wg.Add(1)

	a.initializeSessionClient()
	a.initializeNotificationsClient()
	a.gRPCPostInit()

}

func (a *AdapterServiceClients) InitializeClientsForSession() {
	defer a.wg.Done()

	a.wg.Add(1)

	a.initializeStorageClient()
	a.gRPCPostInit()
}

func (a *AdapterServiceClients) InitializeClientsForStorage() {
	defer a.wg.Done()

	a.wg.Add(1)

	a.initializeSessionClient()
	a.initializeNotificationsClient()
	a.gRPCPostInit()

}

func (a *AdapterServiceClients) StopProtocolListener() {
	for _, v := range a.clientList {
		a.monitor.LogGenericInfo("closing gRPC connection: " + v.Target())
		a.CloseClientConnection(v)
	}

	os.Exit(0)
}
