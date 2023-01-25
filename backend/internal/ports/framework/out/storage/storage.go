package ports

type StoragePort interface {
	CloseConnection()
	Initialize()
}
