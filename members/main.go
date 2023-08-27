package main

import (
	"os"
	"sync"

	apiAdapter "github.com/codeclout/AccountEd/members/adapters/api"
	drivenAdapter "github.com/codeclout/AccountEd/members/adapters/framework/driven"
	memberDriverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers"
	memberProtocolDriverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	"github.com/codeclout/AccountEd/members/ports/framework/driven"
	memberProtocolDriverPort "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols"

	driverAdapterServerConfiguration "github.com/codeclout/AccountEd/members/adapters/framework/drivers/server"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/members/ports/core"
	"github.com/codeclout/AccountEd/members/ports/framework/drivers"

	coreAdapter "github.com/codeclout/AccountEd/members/adapters/core"

	"github.com/codeclout/AccountEd/members/ports/framework/drivers/server"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	serverProtocolAdapter "github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	"github.com/codeclout/AccountEd/pkg/server/server-types/protocols"
)

func main() {
	var (
		homeschoolAPI       api.HomeschoolAPI
		homeschoolCore      core.HomeschoolCore
		homeschoolDriver    drivers.HomeschoolDriverPort
		homeSchoolDriven    driven.HomeschoolDrivenPort
		memberConfiguration server.MembersConfigurationPort
		memberHTTPProtocol  memberProtocolDriverPort.MemberProtocolHTTPPort

		wg sync.WaitGroup
	)

	monitor := monitoring.NewAdapter()

	memberConfiguration = driverAdapterServerConfiguration.NewAdapter(*monitor, "./config.hcl")
	config := *memberConfiguration.LoadMemberConfig()

	applicationName, ok := config["Name"].(string)
	if !ok {
		monitor.LogGenericError("application name not configured")
		os.Exit(1)
	}

	isAppGetOnly, ok := config["GetOnlyConstraint"].(bool)
	if !ok {
		monitor.LogGenericError("is_app_get_only not configured")
		os.Exit(1)
	}

	metadata := protocols.ServerProtocolHttpMetadata{
		RoutePrefix:      "members",
		ServerName:       applicationName,
		UseOnlyGetRoutes: isAppGetOnly,
	}

	httpAdapter := serverProtocolAdapter.NewAdapter(config, metadata, *monitor, &wg)

	memberHTTPProtocolAdapter := memberProtocolDriverAdapter.NewAdapter(config, httpAdapter.FrameworkDriver, *monitor, &wg)
	memberHTTPProtocol = memberHTTPProtocolAdapter
	defer memberHTTPProtocol.StopProtocolListener(httpAdapter.FrameworkDriver)

	gRPCAdapter := serverProtocolAdapter.NewGrpcAdapter(config, *monitor, &wg)
	go gRPCAdapter.InitializeClients()
	defer gRPCAdapter.StopProtocolListener()

	homeschoolCore = coreAdapter.NewAdapter(config, *monitor)
	homeSchoolDriven = drivenAdapter.NewAdapter(*monitor)
	homeschoolAPI = apiAdapter.NewAdapter(config, homeschoolCore, gRPCAdapter, homeSchoolDriven, *monitor)
	homeschoolDriver = memberDriverAdapter.NewAdapter(config, homeschoolAPI, *monitor)

	homeschoolRoutes := homeschoolDriver.InitializeAPI()
	httpAdapter.InitializeRoutes(homeschoolRoutes)

	port, e := httpAdapter.GetPort()
	if e != nil {
		monitor.LogGenericError("member service port: not set")
		panic(e)
	}

	memberHTTPProtocol.Run(port)
}
