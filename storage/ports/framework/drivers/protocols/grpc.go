package protocols

type StorageGrpcDriverPort interface {
	Run()
	InitializeServiceClients(port string)
}
