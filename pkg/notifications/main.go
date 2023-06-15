package main

import (
	"sync"

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

// main initializes and runs the components of the email notification service, including monitoring, configuration, core logic, API, driving,
// and gRPC protocol. It creates instances of the necessary system components and its dependencies, and starts the gRPC server to expose the service.
func main() {
	var (
		notificationConfiguration server.NotificationsConfigurationPort
		emailNotificationDriver   drivers.EmailDriverPort
		emailNotificationApi      apiEmail.EmailApiPort
		emailNotificationCore     core.EmailCorePort
		emailNotificationDriven   driven.EmailDrivenPort
		grpcProtocol              protocols.GRPCProtocolPort

		wg sync.WaitGroup
	)

	monitor := monitoring.NewAdapter()
	wg.Add(1)
	go monitor.Initialize(&wg)

	notificationConfiguration = configuration.NewAdapter(monitor.Logger)
	internalConfig := notificationConfiguration.LoadNotificationsConfig()

	emailNotificationDriven = drivenAdapter.NewAdapter(*internalConfig, monitor.Logger)
	emailNotificationCore = coreAdapter.NewAdapter(*internalConfig, monitor.Logger)
	emailNotificationApi = api.NewAdapter(monitor.Logger, emailNotificationCore, emailNotificationDriven)
	emailNotificationDriver = driverAdapter.NewAdapter(emailNotificationApi, *internalConfig, monitor.Logger)

	grpcProtocol = protocolGrpcAdapter.NewAdapter(*internalConfig, monitor.Logger, emailNotificationDriver)
	grpcProtocol.Run()
}
