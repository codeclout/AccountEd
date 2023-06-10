package protocols

import (
	"os"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type ClientAdapter struct {
	Emailclient    *pb.EmailNotificationServiceClient
	grpcConnection *grpc.ClientConn
	log            *slog.Logger
}

func NewAdapterClientProtocolGRPC(log *slog.Logger) *ClientAdapter {
	return &ClientAdapter{
		log: log,
	}
}

// InitializeClient connects to the email notification service via gRPC. It expects a port string as an input
// and initializes the Emailclient and grpcConnection fields of the ClientAdapter. In case the connection fails, the application
// is terminated with a log error indicating "notifications connection failed".
func (a *ClientAdapter) InitializeClient(port string) {
	connection, e := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // FIXME
	if e != nil {
		a.log.Error("notifications connection failed")
		os.Exit(1)
	}

	emailNotificationClient := pb.NewEmailNotificationServiceClient(connection)

	a.Emailclient = &emailNotificationClient
	a.grpcConnection = connection
}
