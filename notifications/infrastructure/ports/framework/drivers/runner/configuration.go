package runner

type NotificationsInfraConfigPort interface {
	LoadStaticConfig(baseConfiguration map[string]interface{}) *map[string]interface{}
}
