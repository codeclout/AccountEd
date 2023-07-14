package runner

type NotificationsInfraConfigPort interface {
  LoadNotificationsInfrastructureConfig() *map[string]interface{}
}
