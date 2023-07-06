package protocols

type GrpcProtocolPort interface {
	InitializeNotificationsClient(port string)
	InitializeSessionClient(port string)
}
