package protocols

type MemberProtocolGrpcPort interface {
	InitializeClients()
	InitializeClientsForStorage()
}
