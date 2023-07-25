package protocol

import "google.golang.org/grpc"

type ServerProtocolGrpcPort interface {
	CloseClientConnection(conn *grpc.ClientConn)
}
