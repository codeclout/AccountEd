package main

import (
	"fmt"

	acctTypeApiAdapter "github.com/codeclout/AccountEd/adapters/app/api/account-types"
	acctTypeCoreAdapter "github.com/codeclout/AccountEd/adapters/core/account-types"
	hclFramewkInAdapter "github.com/codeclout/AccountEd/adapters/framework/in/hcl"
	httpFramewkInAdapter "github.com/codeclout/AccountEd/adapters/framework/in/http"
	dbFramewkOutAdapter "github.com/codeclout/AccountEd/adapters/framework/out/db"
	"github.com/codeclout/AccountEd/adapters/framework/out/logger"
	acctTypeApiPort "github.com/codeclout/AccountEd/ports/api/account-types"
	acctTypeCorePort "github.com/codeclout/AccountEd/ports/core/account-types"
	hclFramewkInPort "github.com/codeclout/AccountEd/ports/framework/in/hcl"
	httpFramewkInPort "github.com/codeclout/AccountEd/ports/framework/in/http"
	dbFramewkOutPort "github.com/codeclout/AccountEd/ports/framework/out/db"
	loggerFramewkOutPort "github.com/codeclout/AccountEd/ports/framework/out/logger"
)

func main() {
	var (
		e error

		accountTypeDbAdapter   dbFramewkOutPort.UserAccountTypeDbPort
		accountTypeCoreAdapter acctTypeCorePort.UserAccountTypeCorePort
		accountTypeApiAdapter  acctTypeApiPort.UserAccountTypeApiPort
		configAdapter          hclFramewkInPort.RuntimeConfigPort
		httpFrameworkInAdapter httpFramewkInPort.HttpFrameworkInPort
		logFrameworkOutAdapter loggerFramewkOutPort.LogFrameworkOutPort

		configFile = []byte("config.hcl")
	)

	uri := "mongodb://db,db1,db2/accountEd?replicaSet=rs0"

	logFrameworkOutAdapter = logger.NewAdapter()
	go logFrameworkOutAdapter.Initialize()

	configAdapter = hclFramewkInAdapter.NewAdapter(logFrameworkOutAdapter.Log)

	k := configAdapter.GetConfig(configFile)

	accountTypeDbAdapter, e = dbFramewkOutAdapter.NewAdapter(k, logFrameworkOutAdapter.Log, uri)
	if e != nil {
		logFrameworkOutAdapter.Log("fatal", fmt.Sprintf("Failed to instantiate db connection: %v", e))
	}

	accountTypeCoreAdapter = acctTypeCoreAdapter.NewAdapter(logFrameworkOutAdapter.Log)
	accountTypeApiAdapter = acctTypeApiAdapter.NewAdapter(accountTypeCoreAdapter, accountTypeDbAdapter, logFrameworkOutAdapter.Log)
	httpFrameworkInAdapter = httpFramewkInAdapter.NewAdapter(accountTypeApiAdapter, logFrameworkOutAdapter.Log)

	defer logFrameworkOutAdapter.Sync()
	defer accountTypeDbAdapter.CloseConnection()

	logFrameworkOutAdapter.Log("info", "application starting")
	httpFrameworkInAdapter.Run(logFrameworkOutAdapter.HttpMiddlewareLogger)
}
