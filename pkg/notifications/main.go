package main

import (
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/notifications/adapters/api"
	coreAdapter "github.com/codeclout/AccountEd/pkg/notifications/adapters/core"
	drivenAdapter "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/driven"
	driverAdapter "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/drivers"
	protocolGrpcAdapter "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/drivers/protocols"
	configuration "github.com/codeclout/AccountEd/pkg/notifications/adapters/framework/drivers/server"
	apiEmail "github.com/codeclout/AccountEd/pkg/notifications/ports/api"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/core"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/framework/driven"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/framework/drivers"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/framework/drivers/protocols"
	"github.com/codeclout/AccountEd/pkg/notifications/ports/framework/drivers/server"
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
	go monitor.Initialize()

	notificationConfiguration = configuration.NewAdapter(monitor.Logger)
	internalConfig := notificationConfiguration.LoadNotificationsConfig()

	emailNotificationDriven = drivenAdapter.NewAdapter(*internalConfig)
	emailNotificationCore = coreAdapter.NewAdapter(*internalConfig)
	emailNotificationApi = api.NewAdapter(monitor.Logger, emailNotificationCore, emailNotificationDriven)
	emailNotificationDriver = driverAdapter.NewAdapter(emailNotificationApi, *internalConfig, monitor.Logger)

	grpcProtocol = protocolGrpcAdapter.NewAdapter(*internalConfig, monitor.Logger, emailNotificationDriver)
	grpcProtocol.Run()
}
