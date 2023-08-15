package server

type SessionConfigurationPort interface {
	LoadStorageConfig() *map[string]interface{}
}
