package main

import (
	"encoding/json"
	"fmt"

	cloudAdapterIn "github.com/codeclout/AccountEd/adapters/framework/in/aws"
	hclAdapterIn "github.com/codeclout/AccountEd/adapters/framework/in/hcl"
	httpAdapterIn "github.com/codeclout/AccountEd/adapters/framework/in/http"
	cloudPortIn "github.com/codeclout/AccountEd/ports/framework/in/aws"
	hclPortIn "github.com/codeclout/AccountEd/ports/framework/in/hcl"
	httpPortIn "github.com/codeclout/AccountEd/ports/framework/in/http"

	accountTypeAdapterAPI "github.com/codeclout/AccountEd/adapters/api/account-types"
	accountTypeAdapterCore "github.com/codeclout/AccountEd/adapters/core/account-types"
	accountTypeAdapterOut "github.com/codeclout/AccountEd/adapters/framework/out/account-types"

	accountTypePortAPI "github.com/codeclout/AccountEd/ports/api/account-types"
	accountTypePortCore "github.com/codeclout/AccountEd/ports/core/account-types"

	storageAdapterOut "github.com/codeclout/AccountEd/adapters/framework/out/storage"
	logPortOut "github.com/codeclout/AccountEd/ports/framework/out/logger"
	storagePortOut "github.com/codeclout/AccountEd/ports/framework/out/storage"

	"github.com/codeclout/AccountEd/adapters/framework/out/logger"
)

var (
	config map[string]interface{}
)

func main() {
	var (
		e error

		accountTypeCoreAdapter    accountTypePortCore.UserAccountTypeCorePort
		accountTypeApiAdapter     accountTypePortAPI.UserAccountTypeApiPort
		cloudAdapter              cloudPortIn.ParameterPort
		httpInAdapter             httpPortIn.HttpFrameworkInPort
		logAdapterOut             logPortOut.LogFrameworkOutPort
		runtimeConfigAdapter      hclPortIn.RuntimeConfigPort
		storageDefaultAdapter     storagePortOut.StoragePort
		storageAccountTypeAdapter storagePortOut.AccountTypeActionPort

		configFile        = []byte("config.hcl")
		dbconnectionParam string
	)

	logAdapterOut = logger.NewAdapter()

	go logAdapterOut.Initialize()
	defer logAdapterOut.Sync()

	runtimeConfigAdapter = hclAdapterIn.NewAdapter(logAdapterOut.Log)
	rtc := runtimeConfigAdapter.GetConfig(configFile)

	e = json.Unmarshal(rtc, &config)
	if e != nil {
		logAdapterOut.Log("fatal", e.Error())
	}

	cloudAdapter = cloudAdapterIn.NewAdapter(logAdapterOut.Log, config)
	dbconnectionParam = config["DbConnectionParam"].(string)

	uri, e := cloudAdapter.GetSecret(&dbconnectionParam)
	if e != nil {
		logAdapterOut.Log("fatal", fmt.Sprintf("Failed to get db secret: %v", e))
	}

	u, e := cloudAdapter.GetRoleConnectionString(uri)
	if e != nil {
		logAdapterOut.Log("fatal", "unable to retrieve IAM role connection string")
	}

	storageAdapter, e := storageAdapterOut.NewAdapter(rtc, logAdapterOut.Log, u)
	storageDefaultAdapter = storageAdapter

	storageDefaultAdapter.Initialize()
	defer storageDefaultAdapter.CloseConnection()

	storageAccountTypeAdapter = accountTypeAdapterOut.NewAdapter(storageAdapter.GetMongoAccountTypeActions(), logAdapterOut.Log)

	accountTypeCoreAdapter = accountTypeAdapterCore.NewAdapter(logAdapterOut.Log)
	accountTypeApiAdapter = accountTypeAdapterAPI.NewAdapter(accountTypeCoreAdapter, storageAccountTypeAdapter, logAdapterOut.Log)
	httpInAdapter = httpAdapterIn.NewAdapter(accountTypeApiAdapter, logAdapterOut.Log)

	logAdapterOut.Log("info", "application starting")
	httpInAdapter.Run(logAdapterOut.HttpMiddlewareLogger)
}
