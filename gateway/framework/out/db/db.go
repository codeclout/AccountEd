package ports

type DatabasePort interface {
	CloseConnection()
	LogDataInteraction(key string, value string)
}
