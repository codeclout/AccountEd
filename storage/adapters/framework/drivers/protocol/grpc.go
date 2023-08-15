package protocol

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/framework/drivers"
)

type cfg = map[string]interface{}
type dynamosrv = drivers.DynamoDbDriverPort

type Adapter struct {
	awsSessionClient  *awspb.AWSResourceClientServiceClient
	config            cfg
	dynamoDriver      drivers.DynamoDbDriverPort
	listener          net.Listener
	monitor           monitoring.Adapter
	server            *grpc.Server
	waitgroup         *sync.WaitGroup
}

func NewAdapter(config cfg, dynamoDriver dynamosrv, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	return &Adapter{
		config:       config,
		dynamoDriver: dynamoDriver,
		monitor:      monitor,
		waitgroup:    wg,
	}
}

func (a *Adapter) StorageRun() {
	var options []grpc.ServerOption

	port, ok := a.config["Port"].(string)
	if !ok {
		panic("ambiguous port number -> gRPC server")
	}

	listener, e := net.Listen("tcp", ":"+port)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}

	a.listener = listener

	server := grpc.NewServer(options...)
	a.server = server

	dynamov1.RegisterDynamoDBStorageServiceServer(server, a.dynamoDriver)
	reflection.Register(server)

	if e := server.Serve(listener); e != nil {
		a.monitor.LogGenericError(e.Error())
		os.Exit(1)
	}
}

func (a *Adapter) PostInit() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	a.monitor.Logger.Warn("grpc client shutting down")

	a.server.Stop()
	_ = a.listener.Close()

	a.waitgroup.Done()
	os.Exit(0)
}
