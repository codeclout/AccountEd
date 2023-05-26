package main

import (
  "github.com/codeclout/AccountEd/pkg/monitoring"
  "github.com/codeclout/AccountEd/pkg/notifications/adapters/api"
  coreAdapter "github.com/codeclout/AccountEd/pkg/notifications/adapters/core"
  driven2 "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/driven"
  driverAdapter "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/drivers"
  configuration "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/drivers/server"
  apiEmail "github.com/codeclout/AccountEd/pkg/notifications/ports/api"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/core"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/framework/driven"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/framework/drivers"
  "github.com/codeclout/AccountEd/pkg/notifications/ports/framework/drivers/server"
)

func main() {
  var (
    notificationConfiguration server.NotificationsConfigurationPort
    emailNotificationDriver   drivers.EmailDriverPort
    emailNotificationApi      apiEmail.EmailApiPort
    emailNotificationCore     core.EmailCorePort
    emailNotificationDriven   driven.EmailDrivenPort
  )

  notificationConfiguration = configuration.NewAdapter()
  config := notificationConfiguration.LoadNotificationsConfig()

  monitor := monitoring.NewAdapter()
  go monitor.Initialize()

  emailNotificationDriven = driven2.NewAdapter(notificationConfiguration)
  emailNotificationCore = coreAdapter.NewAdapter(emailNotificationDriven)
  emailNotificationApi = api.NewAdapter(monitor.Logger, emailNotificationCore)
  emailNotificationDriver = driverAdapter.NewAdapter(emailNotificationApi, *config, monitor.Logger)
}
