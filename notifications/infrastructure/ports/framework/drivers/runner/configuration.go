package runner

type NotificationsInfraConfigPort interface {
	LoadStorageInfrastructureConfig(baseConfiguration map[string]interface{}) *map[string]interface{}
}
