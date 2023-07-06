package protocols

type httpProtocol interface {
	InitializeNotificationsClient(port string)
}
