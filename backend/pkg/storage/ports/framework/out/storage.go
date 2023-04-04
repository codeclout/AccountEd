package out

type StoragePort interface {
	CloseConnection()
	Initialize()
}
