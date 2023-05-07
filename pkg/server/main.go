package server

import (
  driversAdapter "github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers"
  "github.com/codeclout/AccountEd/pkg/server/ports/framework/drivers"
)

func main() {
  var (
    serverConfigurationAdapter drivers.ConfigurationPort
  )

  configAdapter := driversAdapter.NewAdapter()
}
