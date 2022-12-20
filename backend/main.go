package main

import (
	"encoding/json"
	"fmt"

	acctTypeApiAdapter "github.com/codeclout/AccountEd/adapters/app/api/account-types"
	acctTypeCoreAdapter "github.com/codeclout/AccountEd/adapters/core/account-types"
	cloudFramewkInAdapter "github.com/codeclout/AccountEd/adapters/framework/in/aws"
	hclFramewkInAdapter "github.com/codeclout/AccountEd/adapters/framework/in/hcl"
	httpFramewkInAdapter "github.com/codeclout/AccountEd/adapters/framework/in/http"
	dbFramewkOutAdapter "github.com/codeclout/AccountEd/adapters/framework/out/db"
	"github.com/codeclout/AccountEd/adapters/framework/out/logger"
	acctTypeApiPort "github.com/codeclout/AccountEd/ports/api/account-types"
	acctTypeCorePort "github.com/codeclout/AccountEd/ports/core/account-types"
	cloudFramewkInPort "github.com/codeclout/AccountEd/ports/framework/in/aws"
	hclFramewkInPort "github.com/codeclout/AccountEd/ports/framework/in/hcl"
	httpFramewkInPort "github.com/codeclout/AccountEd/ports/framework/in/http"
	dbFramewkOutPort "github.com/codeclout/AccountEd/ports/framework/out/db"
	loggerFramewkOutPort "github.com/codeclout/AccountEd/ports/framework/out/logger"
)

var (
	config map[string]interface{}
)

func main() {
	var (
		e error

		accountTypeDbAdapter   dbFramewkOutPort.UserAccountTypeDbPort
		accountTypeCoreAdapter acctTypeCorePort.UserAccountTypeCorePort
		accountTypeApiAdapter  acctTypeApiPort.UserAccountTypeApiPort
		cloudAdapter           cloudFramewkInPort.CredentialsPort
		configAdapter          hclFramewkInPort.RuntimeConfigPort
		httpFrameworkInAdapter httpFramewkInPort.HttpFrameworkInPort
		logFrameworkOutAdapter loggerFramewkOutPort.LogFrameworkOutPort
		parameterAdapter       cloudFramewkInPort.ParameterPort

		configFile        = []byte("config.hcl")
		dbconnectionParam string
	)

	logFrameworkOutAdapter = logger.NewAdapter()
	go logFrameworkOutAdapter.Initialize()

	configAdapter = hclFramewkInAdapter.NewAdapter(logFrameworkOutAdapter.Log)
	k := configAdapter.GetConfig(configFile)

	e = json.Unmarshal(k, &config)
	if e != nil {
		logFrameworkOutAdapter.Log("fatal", e.Error())
	}

	cloudAdapter = cloudFramewkInAdapter.NewAdapter(logFrameworkOutAdapter.Log, config)
	c := cloudAdapter.LoadCreds()

	parameterAdapter = cloudFramewkInAdapter.NewAdapter(logFrameworkOutAdapter.Log, config)
	dbconnectionParam = config["DbConnectionParam"].(string)

	uri, e := parameterAdapter.GetSecret(c, &dbconnectionParam)
	if e != nil {
		logFrameworkOutAdapter.Log("fatal", fmt.Sprintf("Failed to get db secret: %v", e))
	}

	accountTypeDbAdapter, e = dbFramewkOutAdapter.NewAdapter(k, logFrameworkOutAdapter.Log, *uri)
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
