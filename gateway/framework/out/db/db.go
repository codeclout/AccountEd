package ports

type DatabasePort interface {
	CloseConnection()
	LogDataInteraction()
}
