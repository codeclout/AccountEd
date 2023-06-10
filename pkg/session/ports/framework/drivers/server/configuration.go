package server

type SessionConfigurationPort interface {
	LoadSessionConfig() *map[string]interface{}
}
