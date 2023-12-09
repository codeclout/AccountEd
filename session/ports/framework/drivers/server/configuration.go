package server

type SessionConfigurationPort interface {
	LoadStaticConfig() *map[string]interface{}
}
