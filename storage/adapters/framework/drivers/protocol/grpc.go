package protocol

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/framework/drivers"
)

type Adapter struct {
	awsSessionClient  *awspb.AWSResourceClientServiceClient
	config            map[string]interface{}
	dynamoDriver      drivers.DynamoDbDriverPort
	monitor           monitoring.Adapter
	sessionConnection *grpc.ClientConn
	waitgroup         *sync.WaitGroup
}

func NewAdapter(config map[string]interface{}, dynamoDriver drivers.DynamoDbDriverPort, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	return &Adapter{
		config:       config,
		dynamoDriver: dynamoDriver,
		monitor:      monitor,
		waitgroup:    wg,
	}
}

func (a *Adapter) Run() {
	var options []grpc.ServerOption

	listener, e := net.Listen("tcp", ":9092")
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}

	server := grpc.NewServer(options...)
	dynamov1.RegisterDynamoDBStorageServiceServer(server, a.dynamoDriver)
	reflection.Register(server)

	if e := server.Serve(listener); e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}
}

func (a *Adapter) InitializeServiceClients(port string) {
	connection, e := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // @TODO
	if e != nil {
		a.monitor.LogGenericError("session connection failed")
		os.Exit(1)
	}

	awsSessionClient := awspb.NewAWSResourceClientServiceClient(connection)
	a.awsSessionClient = &awsSessionClient
}

func (a *Adapter) PostInit() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	a.monitor.Logger.Warn("grpc client shutting down")

	a.waitgroup.Done()
	os.Exit(0)
}
