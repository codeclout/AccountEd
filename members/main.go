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

	applicationName, ok := config["application_name"]
	if !ok {
		monitor.Logger.Error("application_name not configured")
		os.Exit(1)
	}

	routePrefix, ok := config["route_prefix"]
	if !ok {
		monitor.Logger.Error("route_prefix not configured")
		os.Exit(1)
	}

	isAppGetOnly, ok := config["is_app_get_only"]
	if !ok {
		monitor.Logger.Error("is_app_get_only not configured")
		os.Exit(1)
	}

	protocolAdapter := protocol.NewAdapter(
		applicationName.(string),
		routePrefix.(string),
		isAppGetOnly.(bool),
		monitor.Logger,
		monitor.HttpMiddlewareLogger,
		&wg)

	protocolDriver = protocolAdapter

	grpcClientAdapter := driverGrpcAdapter.NewAdapterClientProtocolGRPC(monitor.Logger)
	grpcClientAdapter.InitializeClient("9000")

	homeschoolCore = coreAdapter.NewAdapter(monitor.Logger)
	homeschoolAPI = apiAdapter.NewAdapter(homeschoolCore, grpcClientAdapter, monitor.Logger)
	homeschoolDriver = driverAdapter.NewAdapter(homeschoolAPI, monitor.Logger, config)
	homeschoolRoutes := homeschoolDriver.InitializeAPI(protocolAdapter.HTTP)

	// sessionDriver, _ = driverAdapterSession.NewAdapter(monitor)
	http, port := protocolDriver.Initialize(homeschoolRoutes)

	wg.Add(1)
	go protocolAdapter.PostInit(http, &wg)
	defer protocolAdapter.StopProtocolListener(http)

	app := driverAdapterProtocol.NewAdapter(monitor.Logger, http, httpmiddleware.NewLoggerMiddleware)
	app.InitializeClient(port)
}
