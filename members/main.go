package members

import (
  apiAdapter "github.com/codeclout/AccountEd/members/adapters/api"
  coreAdapter "github.com/codeclout/AccountEd/members/adapters/core"
  driverAdapter "github.com/codeclout/AccountEd/members/adapters/framework/drivers"
  driverAdapterServerConfiguration "github.com/codeclout/AccountEd/members/adapters/framework/drivers/server"
  "github.com/codeclout/AccountEd/members/ports/api"
  "github.com/codeclout/AccountEd/members/ports/core"
  "github.com/codeclout/AccountEd/members/ports/framework/drivers"
  "github.com/codeclout/AccountEd/members/ports/framework/drivers/server"
  "github.com/codeclout/AccountEd/pkg/monitoring"
  "github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
  "github.com/codeclout/AccountEd/pkg/server/ports/framework/drivers/protocols"
  driverAdapterSession "github.com/codeclout/AccountEd/pkg/session/adapter/framework/driver"
  "github.com/codeclout/AccountEd/pkg/session/ports/framework/driver"
)

func main() {
  var (
    homeschoolApi       api.HomeschoolAPI
    homeschoolCore      core.HomeschoolCore
    homeschooDriver     drivers.HomeschoolDriverPort
    memberConfiguration server.MembersConfigurationPort
    protocolDriver      protocols.ProtocolPort
    sessionDriver       driver.SessionPort
  )

  monitor := monitoring.NewAdapter()
  go monitor.Initialize()

  monitor.Logger.Info("starting application ->", monitor.GetTimeStamp())

  memberConfiguration = driverAdapterServerConfiguration.NewAdapter(monitor)
  cfg := *memberConfiguration.LoadMemberConfig()

  protocolAdapter := protocol.NewAdapter(
    cfg["applicationName"].(string),
    cfg["routePrefix"].(string),
    cfg["isAppGetOnly"].(bool),
    monitor)

  protocolDriver = protocolAdapter

  homeschoolApi = apiAdapter.NewAdapter(homeschoolCore, monitor)
  homeschooDriver = driverAdapter.NewAdapter(homeschoolApi, monitor)
  homeschoolCore = coreAdapter.NewAdapter(monitor)

  homeschoolRoutes := homeschooDriver.InitializeAPI(protocolAdapter.HTTP)

  sessionDriver, _ = driverAdapterSession.NewAdapter(monitor)

  http := protocolDriver.Initialize(homeschoolRoutes)

}
