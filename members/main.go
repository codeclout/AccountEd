package main

import (
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"

	apiAdapter "github.com/codeclout/AccountEd/members/adapters/api"
	drivenAdapter "github.com/codeclout/AccountEd/members/adapters/framework/driven"
	memberDriverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers"
	memberProtocolDriverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers/protocols"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/members/ports/framework/driven"
	memberProtocolDriverPort "github.com/codeclout/AccountEd/members/ports/framework/drivers/protocols"

	driverAdapterServerConfiguration "github.com/codeclout/AccountEd/members/adapters/framework/drivers/server"
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
		memberAPI        api.MemberAPI
		memberCore       core.MemberCorePort
		homeschoolDriver drivers.HomeschoolDriverPort
		memberDrivenPort driven.MemberDrivenPort

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

	memberCore = coreAdapter.NewAdapter(config, *monitor)
	memberDrivenPort = drivenAdapter.NewAdapter(*monitor)
	memberAPI = apiAdapter.NewAdapter(config, memberCore, gRPCAdapter, memberDrivenPort, *monitor)

	homeschoolDriver = memberDriverAdapter.NewHomeschoolAdapter(config, memberAPI, *monitor)
	homeschoolMemberRoutes := homeschoolDriver.InitializeHomeschoolAPI()

	memberDriver := memberDriverAdapter.NewAdapter(config, memberAPI, *monitor)
	memberRoutes := memberDriver.InitializeMemberAPI()

	var routes []*fiber.App
	var r = append(homeschoolMemberRoutes, memberRoutes...)

	httpAdapter.InitializeRoutes(append(routes, r...))

	port, e := httpAdapter.GetPort()
	if e != nil {
		monitor.LogGenericError("member service port: not set")
		panic(e)
	}

	memberHTTPProtocol.Run(port)
}
