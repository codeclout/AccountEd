package main

import (
	apiAdapter "github.com/codeclout/AccountEd/members/adapters/api"
	coreAdapter "github.com/codeclout/AccountEd/members/adapters/core"
	driverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers"
	driverAdapterProtocolHTTP "github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	driverAdapterServerConfiguration "github.com/codeclout/AccountEd/members/adapters/framework/drivers/server"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/members/ports/core"
	"github.com/codeclout/AccountEd/members/ports/framework/drivers"
	httpmiddleware "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols/http-middleware"
	"github.com/codeclout/AccountEd/members/ports/framework/drivers/server"
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	"github.com/codeclout/AccountEd/pkg/server/ports/framework/drivers/protocols"
)

func main() {
	var (
		homeschoolAPI       api.HomeschoolAPI
		homeschoolCore      core.HomeschoolCore
		homeschoolDriver    drivers.HomeschoolDriverPort
		memberConfiguration server.MembersConfigurationPort
		protocolDriver      protocols.ProtocolPort
		// sessionDriver       driver.SessionPort
	)

	monitor := monitoring.NewAdapter()
	go monitor.Initialize()

	memberConfiguration = driverAdapterServerConfiguration.NewAdapter(monitor.Logger)
	config := *memberConfiguration.LoadMemberConfig()

	protocolAdapter := protocol.NewAdapter(
		config["application_name"].(string),
		config["route_prefix"].(string),
		config["is_app_get_only"].(bool),
		monitor.Logger)

	protocolDriver = protocolAdapter

	homeschoolCore = coreAdapter.NewAdapter(monitor.Logger)
	homeschoolAPI = apiAdapter.NewAdapter(homeschoolCore, monitor.Logger)
	homeschoolDriver = driverAdapter.NewAdapter(homeschoolAPI, monitor.Logger, config)
	homeschoolRoutes := homeschoolDriver.InitializeAPI(protocolAdapter.HTTP)

	// sessionDriver, _ = driverAdapterSession.NewAdapter(monitor)
	http := protocolDriver.Initialize(homeschoolRoutes)
	go protocolAdapter.PostInit(http)

	app := driverAdapterProtocolHTTP.NewAdapter(monitor.Logger, http, httpmiddleware.NewLoggerMiddleware)

	app.Run(monitor.HttpMiddlewareLogger)

}
