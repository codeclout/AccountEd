package main

import (
	"encoding/json"
	"fmt"

	accountTypeAdapterAPI "github.com/codeclout/AccountEd/internal/adapters/api/account-types"
	postalCodeAdapterAPI "github.com/codeclout/AccountEd/internal/adapters/api/postal-codes"
	accountTypeAdapterCore "github.com/codeclout/AccountEd/internal/adapters/core/account-types"
	postalCodeAdapterCore "github.com/codeclout/AccountEd/internal/adapters/core/postal-codes"
	cloudAdapterIn "github.com/codeclout/AccountEd/internal/adapters/framework/in/aws"
	mapAdapterIn "github.com/codeclout/AccountEd/internal/adapters/framework/in/gcp"
	hclAdapterIn "github.com/codeclout/AccountEd/internal/adapters/framework/in/hcl"
	httpAdapterIn "github.com/codeclout/AccountEd/internal/adapters/framework/in/http"
	accountTypeAdapterOut "github.com/codeclout/AccountEd/internal/adapters/framework/out/account-types"
	"github.com/codeclout/AccountEd/internal/adapters/framework/out/logger"
	storageAdapterOut "github.com/codeclout/AccountEd/internal/adapters/framework/out/storage"
	accountTypePortAPI "github.com/codeclout/AccountEd/internal/ports/api/account-types"
	postalCodePortAPI "github.com/codeclout/AccountEd/internal/ports/api/postal-codes"
	accountTypePortCore "github.com/codeclout/AccountEd/internal/ports/core/account-types"
	postalCodePortCore "github.com/codeclout/AccountEd/internal/ports/core/postal-codes"
	cloudPortIn "github.com/codeclout/AccountEd/internal/ports/framework/in/aws"
	mapPortIn "github.com/codeclout/AccountEd/internal/ports/framework/in/gcp"
	hclPortIn "github.com/codeclout/AccountEd/internal/ports/framework/in/hcl"
	httpPortIn "github.com/codeclout/AccountEd/internal/ports/framework/in/http"
	logPortOut "github.com/codeclout/AccountEd/internal/ports/framework/out/logger"
	"github.com/codeclout/AccountEd/internal/ports/framework/out/storage"
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
		mapAdapter                mapPortIn.PostalCodeFrameworkIn
		postalCodeCoreAdapter     postalCodePortCore.PostalCodeCorePort
		postalCodeApiAdapter      postalCodePortAPI.PostalCodeApiPort
		runtimeConfigAdapter      hclPortIn.RuntimeConfigPort
		storageDefaultAdapter     ports.StoragePort
		storageAccountTypeAdapter ports.AccountTypeActionPort

		configFile        = []byte("./config.hcl")
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

	mapAdapter = mapAdapterIn.NewAdapter(config, logAdapterOut.Log)
	postalCodeCoreAdapter = postalCodeAdapterCore.NewAdapter(logAdapterOut.Log)
	postalCodeApiAdapter = postalCodeAdapterAPI.NewAdapter(postalCodeCoreAdapter, mapAdapter, logAdapterOut.Log)

	httpInAdapter = httpAdapterIn.NewAdapter(accountTypeApiAdapter, postalCodeApiAdapter, logAdapterOut.Log)

	logAdapterOut.Log("info", "application starting")
	httpInAdapter.Run(logAdapterOut.HttpMiddlewareLogger)
}
