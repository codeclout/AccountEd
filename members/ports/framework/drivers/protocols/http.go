package protocols

type httpProtocol interface {
	InitializeClient(port string)
}
