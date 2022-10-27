package main

import (
	"fmt"

	account2 "github.com/codeclout/AccountEd/adapters/app/api/account"
	"github.com/codeclout/AccountEd/adapters/core/account"
	"github.com/codeclout/AccountEd/adapters/framework/in/http"
	"github.com/codeclout/AccountEd/adapters/framework/out/db"
	"github.com/codeclout/AccountEd/adapters/framework/out/logger"
	accountPort "github.com/codeclout/AccountEd/ports/app"
	corePort "github.com/codeclout/AccountEd/ports/core"
	httpPort "github.com/codeclout/AccountEd/ports/framework/in/http"
	accountDbPort "github.com/codeclout/AccountEd/ports/framework/out/db"
	loggerPort "github.com/codeclout/AccountEd/ports/framework/out/logger"
)

func main() {
	var (
		e error

		accountDbAdapter accountDbPort.AccountDbPort
		accountAdapter   corePort.AccountPort
		accountAPI       accountPort.AccountAPIPort
		httpAdapter      httpPort.HTTPPort
		loggerAdapter    loggerPort.LoggerPort
	)

	uri := "mongodb://db,db1,db2/accountEd?replicaSet=rs0"

	accountAdapter = account.NewAdapter()
	loggerAdapter = logger.NewAdapter()

	go loggerAdapter.Initialize()

	accountDbAdapter, e = db.NewAdapter(5, loggerAdapter.Log, uri)
	if e != nil {
		loggerAdapter.Log("fatal", fmt.Sprintf("Failed to instantiate db connection: %v", e))
	}

	accountAPI = account2.NewAdapter(accountAdapter, accountDbAdapter)
	httpAdapter = http.NewAdapter(accountAPI, loggerAdapter.Log)

	defer loggerAdapter.Sync()
	defer accountDbAdapter.CloseConnection()

	loggerAdapter.Log("info", "application starting")
	httpAdapter.Run(loggerAdapter.HttpMiddlewareLogger)
}
