package protocols

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	spb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type ClientAdapter struct {
	AWSClient               *spb.AWSResourceClientServiceClient
	Emailclient             *npb.EmailNotificationServiceClient
	MemberClient            *spb.MemberSessionClient
	log                     *slog.Logger
	notificationsConnection *grpc.ClientConn
	CloudSessionConnection  *grpc.ClientConn
	waitgroup               *sync.WaitGroup
}

func NewAdapterClientProtocolGRPC(log *slog.Logger, wg *sync.WaitGroup) *ClientAdapter {
	return &ClientAdapter{
		log:       log,
		waitgroup: wg,
	}
}

// InitializeNotificationsClient sets up a gRPC connection to the email notification service, and configures its client attribute. This
// function takes a port number as a parameter, creates a gRPC connection to the specified port with insecure transport credentials, and instantiates
// a new EmailNotificationServiceClient. If the connection fails, this function will log an error and cause the program to exit with a non-zero value.
// Upon successful setup, the Emailclient attribute of the ClientAdapter will be set to the new client instance, and the notificationsConnection
// attribute will store the newly created connection.
func (a *ClientAdapter) InitializeNotificationsClient(port string) {
	connection, e := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // @TODO
	if e != nil {
		a.log.Error("notifications connection failed")
		os.Exit(1)
	}

	emailNotificationClient := npb.NewEmailNotificationServiceClient(connection)

	a.Emailclient = &emailNotificationClient
	a.notificationsConnection = connection
}

// InitializeSessionClient sets up a gRPC connection to the AWS resource client service and configures its client attribute.
// This function takes a port number as a parameter, creates a gRPC connection to the specified port with insecure transport credentials, and
// instantiates a new AWSResourceClientServiceClient. If the connection fails, this function logs an error and exits with a non-zero value.
// Upon successful setup, the AWSClient attribute of the ClientAdapter will be set to the new client instance, and the CloudSessionConnection
// attribute will store the newly created connection.
func (a *ClientAdapter) InitializeSessionClient(port string) {
	connection, e := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // @TODO
	if e != nil {
		a.log.Error("session connection failed")
		os.Exit(1)
	}

	awsSessionClient := spb.NewAWSResourceClientServiceClient(connection)
	memberSessionClient := spb.NewMemberSessionClient(connection)

	a.CloudSessionConnection = connection
	a.AWSClient = &awsSessionClient
	a.MemberClient = &memberSessionClient
}

func (a *ClientAdapter) PostInit(wg *sync.WaitGroup) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	<-s
	a.log.Warn("grpc client shutting down")

	wg.Done()
	os.Exit(0)
}

func (a *ClientAdapter) StopProtocolListener() {
	var e error
	a.waitgroup.Wait()

	e = a.CloudSessionConnection.Close()
	if e != nil {
		a.log.Error(e.Error())
	}

	e = a.notificationsConnection.Close()
	if e != nil {
		a.log.Error(e.Error())
	}
}
