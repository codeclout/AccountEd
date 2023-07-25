package main

import (
	"sync"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/storage/adapters/api"
	"github.com/codeclout/AccountEd/storage/adapters/core"
	"github.com/codeclout/AccountEd/storage/adapters/framework/driven"
	storageDriverAdapter "github.com/codeclout/AccountEd/storage/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/storage/adapters/framework/drivers/server"
	"github.com/codeclout/AccountEd/storage/ports/framework/drivers"
)

var staticConfigurationPath = "./config.hcl"

func main() {
	var wg sync.WaitGroup
	var (
		driverAdapter drivers.DynamoDbDriverPort
	)

	monitor := monitoring.NewAdapter()

	storageConfiguration := server.NewAdapter(*monitor, staticConfigurationPath)
	config := *storageConfiguration.LoadSessionConfig()

	drivenAdapter := driven.NewAdapter(config, *monitor, &wg)
	coreAdapter := core.NewAdapter(config, *monitor)
	apiAdapter := api.NewAdapter(config, coreAdapter, *monitor)
	driverAdapterStorage := storageDriverAdapter.NewAdapter(config, apiAdapter, *monitor, &wg)
}
