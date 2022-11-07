package main

import (
	"fmt"

	accountApiAdapter "github.com/codeclout/AccountEd/adapters/app/api/account"
	accountCoreAdapter "github.com/codeclout/AccountEd/adapters/core/account"
	hclFrameworkAdapter "github.com/codeclout/AccountEd/adapters/framework/in/hcl"
	httpFrameworkInAdapter "github.com/codeclout/AccountEd/adapters/framework/in/http"
	dbFrameworkOut "github.com/codeclout/AccountEd/adapters/framework/out/db"
	"github.com/codeclout/AccountEd/adapters/framework/out/logger"
	accountApiPort "github.com/codeclout/AccountEd/ports/app"
	corePort "github.com/codeclout/AccountEd/ports/core"
	httpFrameworkInPort "github.com/codeclout/AccountEd/ports/framework/in/http"
	dbFrameworkOutPort "github.com/codeclout/AccountEd/ports/framework/out/db"
	loggerPort "github.com/codeclout/AccountEd/ports/framework/out/logger"
)

func main() {
	var (
		e error

		accountDbAdapter     dbFrameworkOutPort.AccountDbPort
		accountAdapter       corePort.AccountPort
		accountAPI           accountApiPort.AccountAPIPort
		httpFrameworkAdapter httpFrameworkInPort.HTTPPort
		loggerAdapter        loggerPort.LoggerPort

		configFile = []byte("config.hcl")
	)

	uri := "mongodb://db,db1,db2/accountEd?replicaSet=rs0"

	loggerAdapter = logger.NewAdapter()
	go loggerAdapter.Initialize()

	configAdapter := hclFrameworkAdapter.NewAdapter(loggerAdapter.Log)
	k, e := configAdapter.GetConfig(configFile)
	if e != nil {
		loggerAdapter.Log("fatal", fmt.Sprintf("Failed to get runtime config: %v", e))
	}

	accountDbAdapter, e = dbFrameworkOut.NewAdapter(k, loggerAdapter.Log, uri)
	if e != nil {
		loggerAdapter.Log("fatal", fmt.Sprintf("Failed to instantiate db connection: %v", e))
	}

	accountAdapter = accountCoreAdapter.NewAdapter(loggerAdapter.Log)
	accountAPI = accountApiAdapter.NewAdapter(accountAdapter, accountDbAdapter, loggerAdapter.Log)
	httpFrameworkAdapter = httpFrameworkInAdapter.NewAdapter(accountAPI, loggerAdapter.Log)

	defer loggerAdapter.Sync()
	defer accountDbAdapter.CloseConnection()

	loggerAdapter.Log("info", "application starting")
	httpFrameworkAdapter.Run(loggerAdapter.HttpMiddlewareLogger)
}
