package main

import (
	"os"
	"sync"

	apiAdapter "github.com/codeclout/AccountEd/members/adapters/api"
	coreAdapter "github.com/codeclout/AccountEd/members/adapters/core"
	driverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers"
	driverAdapterProtocol "github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	driverGrpcAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	driverAdapterServerConfiguration "github.com/codeclout/AccountEd/members/adapters/framework/drivers/server"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/members/ports/core"
	"github.com/codeclout/AccountEd/members/ports/framework/drivers"
	httpmiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
	"github.com/codeclout/AccountEd/members/ports/framework/drivers/server"
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	serverProtocols "github.com/codeclout/AccountEd/pkg/server/ports/framework/drivers/protocols"
)

// main initializes the application and connects all dependencies, setting up the server and starting required components.
func main() {
	var (
		homeschoolAPI       api.HomeschoolAPI
		homeschoolCore      core.HomeschoolCore
		homeschoolDriver    drivers.HomeschoolDriverPort
		memberConfiguration server.MembersConfigurationPort
		protocolDriver      serverProtocols.ProtocolPort

		wg sync.WaitGroup
		// sessionDriver       driver.SessionPort
	)

	monitor := monitoring.NewAdapter()
	wg.Add(1)
	go monitor.Initialize(&wg)

	memberConfiguration = driverAdapterServerConfiguration.NewAdapter(monitor.Logger, "./config.hcl")
	config := *memberConfiguration.LoadMemberConfig()

	applicationName, ok := config["Name"].(string)
	if !ok {
		monitor.Logger.Error("application name not configured")
		os.Exit(1)
	}

	routePrefix, ok := config["RoutePrefix"].(string)
	if !ok {
		monitor.Logger.Error("route prefix not configured")
		os.Exit(1)
	}

	isAppGetOnly, ok := config["GetOnlyConstraint"].(bool)
	if !ok {
		monitor.Logger.Error("is_app_get_only not configured")
		os.Exit(1)
	}

	protocolAdapter := protocol.NewAdapter(
		config,
		routePrefix,
		applicationName,
		monitor.Logger,
		monitor.HttpMiddlewareLogger,
		&wg,
		isAppGetOnly)

	protocolDriver = protocolAdapter

	grpcClientAdapter := driverGrpcAdapter.NewAdapterClientProtocolGRPC(monitor.Logger, &wg)

	notificationServicePort, ok := config["NotificationsServicePort"].(string)
	if !ok {
		monitor.Logger.Error("notifications service port not configured")
		os.Exit(1)
	}
	grpcClientAdapter.InitializeNotificationsClient(notificationServicePort)

	sessionServicePort, ok := config["SessionServicePort"].(string)
	if !ok {
		monitor.Logger.Error("session service port not configured")
		os.Exit(1)
	}
	grpcClientAdapter.InitializeSessionClient(sessionServicePort)

	wg.Add(1)
	go grpcClientAdapter.PostInit(&wg)
	defer grpcClientAdapter.StopProtocolListener()

	homeschoolCore = coreAdapter.NewAdapter(config, monitor.Logger)
	homeschoolAPI = apiAdapter.NewAdapter(config, homeschoolCore, grpcClientAdapter, monitor.Logger)
	homeschoolDriver = driverAdapter.NewAdapter(config, homeschoolAPI, monitor.Logger)
	homeschoolRoutes := homeschoolDriver.InitializeAPI(protocolAdapter.HTTP)

	// sessionDriver, _ = driverAdapterSession.NewAdapter(monitor)
	http, port := protocolDriver.Initialize(homeschoolRoutes)

	wg.Add(1)
	go protocolAdapter.PostInit(&wg)
	defer protocolAdapter.StopProtocolListener(http)

	app := driverAdapterProtocol.NewAdapter(config, http, httpmiddleware.NewLoggerMiddleware, monitor.Logger)
	app.InitializeNotificationsClient(port)
}
