package drivers

type ConfigurationPort interface {
	LoadServerConfiguration() *map[string]interface{}
}
