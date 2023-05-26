package server

type NotificationsConfigurationPort interface {
  LoadNotificationsConfig() *map[string]interface{}
}
