package main

import (
	coreAdapter "github.com/codeclout/AccountEd/notifications/adapters/core"
	drivenAdapter "github.com/codeclout/AccountEd/notifications/adapters/framework/driven"
	driverAdapter "github.com/codeclout/AccountEd/notifications/adapters/framework/drivers"
	protocolGrpcAdapter "github.com/codeclout/AccountEd/notifications/adapters/framework/drivers/protocols"
	configuration "github.com/codeclout/AccountEd/notifications/adapters/framework/drivers/server"
	"github.com/codeclout/AccountEd/notifications/ports/core"
	"github.com/codeclout/AccountEd/notifications/ports/framework/driven"
	"github.com/codeclout/AccountEd/notifications/ports/framework/drivers"
	"github.com/codeclout/AccountEd/notifications/ports/framework/drivers/protocols"
	"github.com/codeclout/AccountEd/notifications/ports/framework/drivers/server"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"

	apiEmail "github.com/codeclout/AccountEd/notifications/ports/api"

	"github.com/codeclout/AccountEd/notifications/adapters/api"
)

func main() {
	var (
		notificationConfiguration server.NotificationsConfigurationPort
		emailNotificationDriver   drivers.EmailDriverPort
		emailNotificationApi      apiEmail.EmailApiPort
		emailNotificationCore     core.EmailCorePort
		emailNotificationDriven   driven.EmailDrivenPort
		grpcProtocol              protocols.GRPCProtocolPort
	)

	monitor := monitoring.NewAdapter()

	notificationConfiguration = configuration.NewAdapter(*monitor)
	config := *notificationConfiguration.LoadNotificationsConfig()

	emailNotificationDriven = drivenAdapter.NewAdapter(config, *monitor)
	emailNotificationCore = coreAdapter.NewAdapter(config, *monitor)
	emailNotificationApi = api.NewAdapter(config, emailNotificationCore, emailNotificationDriven, *monitor)
	emailNotificationDriver = driverAdapter.NewAdapter(emailNotificationApi, config, *monitor)

	grpcProtocol = protocolGrpcAdapter.NewAdapter(config, emailNotificationDriver, *monitor)
	grpcProtocol.Run()
}
