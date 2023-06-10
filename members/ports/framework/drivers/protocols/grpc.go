package protocols

type GrpcProtocolPort interface {
	InitializeClient(port string)
}
