package protocols

type MemberProtocolGrpcPort interface {
	InitializeClientsForMember()
	InitializeClientsForSession()
	InitializeClientsForStorage()
}
