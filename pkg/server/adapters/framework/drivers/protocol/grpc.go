package protocol

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	notifypb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	memberpb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type AdapterGrpc struct {
	AWS_SessionClient        *awspb.AWSResourceClientServiceClient
	Email_NotificationClient *notifypb.EmailNotificationServiceClient
	Member_SessionClient     *memberpb.MemberSessionClient
	config                   map[string]interface{}
	clientList               []*grpc.ClientConn
	monitor                  monitoring.Adapter
	wg                       *sync.WaitGroup
}

func NewGrpcAdapter(config map[string]interface{}, monitor monitoring.Adapter, wg *sync.WaitGroup) *AdapterGrpc {
	wg.Add(1)
	go gRPCPostInit(monitor, wg, "closing grpc client connections")

	return &AdapterGrpc{
		config:  config,
		monitor: monitor,
		wg:      wg,
	}
}

func gRPCPostInit(monitor monitoring.Adapter, wg *sync.WaitGroup, warning string) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	monitor.Logger.Warn(warning)

	wg.Done()
	os.Exit(0)
}

func (a *AdapterGrpc) CloseClientConnection(conn *grpc.ClientConn) {
	_ = conn.Close()
}

func (a *AdapterGrpc) InitializeClients() {
	a.wg.Add(1)

	a.initializeSessionClient()
	a.initializeNotificationsClient()

	a.wg.Done()
}

func (a *AdapterGrpc) initializeSessionClient() {
	connection, e := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		a.monitor.LogGenericError("gRPC session client failed to connect")
		os.Exit(1)
	}

	_ = append(a.clientList, connection)
	awsSessionClient := awspb.NewAWSResourceClientServiceClient(connection)
	memberSessionClient := memberpb.NewMemberSessionClient(connection)

	a.AWS_SessionClient = &awsSessionClient
	a.Member_SessionClient = &memberSessionClient
}

func (a *AdapterGrpc) initializeNotificationsClient() {
	connection, e := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		a.monitor.LogGenericError("gRPC notifications client failed to connect")
		os.Exit(1)
	}

	_ = append(a.clientList, connection)
	emailNotificationsClient := notifypb.NewEmailNotificationServiceClient(connection)

	a.Email_NotificationClient = &emailNotificationsClient
}

func (a *AdapterGrpc) StopProtocolListener() {
	for _, v := range a.clientList {
		a.monitor.LogGenericInfo("closing gRPC connection: " + v.Target())
		a.CloseClientConnection(v)
	}
}
