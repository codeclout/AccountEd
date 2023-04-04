package main

import (
	"encoding/json"
	"fmt"

	postalCodeAdapterAPI "github.com/codeclout/AccountEd/internal/adapters/api/postal-codes"
	postalCodeAdapterCore "github.com/codeclout/AccountEd/internal/adapters/core/postal-codes"
	mapAdapterIn "github.com/codeclout/AccountEd/internal/adapters/framework/in/gcp"
	postalCodePortAPI "github.com/codeclout/AccountEd/internal/ports/api/postal-codes"
	postalCodePortCore "github.com/codeclout/AccountEd/internal/ports/core/postal-codes"
	mapPortIn "github.com/codeclout/AccountEd/internal/ports/framework/in/gcp"
	accountTypeAdapterAPI "github.com/codeclout/AccountEd/onboarding/internal/adapters/api/account-types"
	onboardingApiAdapter "github.com/codeclout/AccountEd/onboarding/internal/adapters/api/workflows"
	accountTypeAdapterCore "github.com/codeclout/AccountEd/onboarding/internal/adapters/core/account-types"
	onboardingCoreAdapter "github.com/codeclout/AccountEd/onboarding/internal/adapters/core/workflows"
	httpAdapterIn "github.com/codeclout/AccountEd/onboarding/internal/adapters/framework/in/http"
	frameworkOutStorage "github.com/codeclout/AccountEd/onboarding/internal/adapters/framework/out/storage"
	accountTypePortAPI "github.com/codeclout/AccountEd/onboarding/internal/ports/api/account-types"
	onboardingApiPort "github.com/codeclout/AccountEd/onboarding/internal/ports/api/workflows"
	accountTypePortCore "github.com/codeclout/AccountEd/onboarding/internal/ports/core/account-types"
	"github.com/codeclout/AccountEd/onboarding/internal/ports/core/workflows"
	hclPortIn "github.com/codeclout/AccountEd/onboarding/internal/ports/framework/in/hcl"
	httpPortIn "github.com/codeclout/AccountEd/onboarding/internal/ports/framework/in/http"
	"github.com/codeclout/AccountEd/onboarding/internal/ports/framework/out/storage"
	service_config "github.com/codeclout/AccountEd/onboarding/service-config"
	"github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/out/logger"
	logPortOut "github.com/codeclout/AccountEd/pkg/monitoring/ports/framework/out/logger"
	cloudAdapterIn "github.com/codeclout/AccountEd/pkg/service-identity/adapters/framework/in"
	"github.com/codeclout/AccountEd/pkg/service-identity/ports/framework/in"
	storageAdapterOut "github.com/codeclout/AccountEd/pkg/storage/adapters/framework/out"
	"github.com/codeclout/AccountEd/pkg/storage/ports/framework/out"
)

var (
	config map[string]interface{}
)

func main() {
	var (
		e error

		accountTypeCoreAdapter       accountTypePortCore.UserAccountTypeCorePort
		accountTypeApiAdapter        accountTypePortAPI.UserAccountTypeApiPort
		cloudCredentialsAdapter      in.CredentialsPort
		cloudParamsAdapter           in.ParameterPort
		httpInAdapter                httpPortIn.ServerFrameworkPort
		logAdapterOut                logPortOut.LogFrameworkOutPort
		mapAdapter                   mapPortIn.PostalCodeFrameworkIn
		postalCodeCoreAdapter        postalCodePortCore.PostalCodeCorePort
		postalCodeApiAdapter         postalCodePortAPI.PostalCodeApiPort
		runtimeConfigAdapter         hclPortIn.RuntimeConfigPort
		storageDefaultAdapter        out.StoragePort
		storageAccountTypeOutAdapter storage.AccountTypeActionPort
		storageHomeSchoolAdapter     storage.HomeschoolActionPort
		homeschoolOnboardCoreAdapter workflows.HomeSchoolCorePort
		homeschoolOnboardApiAdapter  onboardingApiPort.OnboardHomeschoolApiPort

		configFile      = []byte("./config.hcl")
		connectionParam string
	)

	logAdapterOut = logger.NewAdapter()

	go logAdapterOut.Initialize()
	defer logAdapterOut.Sync()

	runtimeConfigAdapter = service_config.NewAdapter(logAdapterOut.Log)
	rtc := runtimeConfigAdapter.GetConfig(configFile)

	e = json.Unmarshal(rtc, &config)
	if e != nil {
		logAdapterOut.Log("fatal", e.Error())
	}

	cloudSession := cloudAdapterIn.NewAdapter(logAdapterOut.Log, config)
	cloudCredentialsAdapter = cloudSession
	cloudParamsAdapter = cloudSession

	connectionParam = config["DbConnectionParam"].(string)

	uri, e := cloudParamsAdapter.GetSecret(&connectionParam)
	if e != nil {
		logAdapterOut.Log("fatal", fmt.Sprintf("Failed to get db secret: %v", e))
	}

	atlasConnectionString, e := cloudParamsAdapter.GetRoleConnectionString(uri)
	if e != nil {
		logAdapterOut.Log("fatal", "unable to retrieve IAM role connection string")
	}

	serviceIdentity := cloudCredentialsAdapter.ExportCreds()

	storageAdapter, e := storageAdapterOut.NewAdapter(
		rtc,
		logAdapterOut.Log,
		atlasConnectionString,
		serviceIdentity)

	storageDefaultAdapter = storageAdapter

	storageDefaultAdapter.Initialize()
	defer storageDefaultAdapter.CloseConnection()

	storageAccountTypeOutAdapter = frameworkOutStorage.NewAdapter(
		storageAdapter.GetMongoActions(),
		logAdapterOut.Log)

	storageHomeSchoolAdapter = frameworkOutStorage.NewAdapter(
		storageAdapter.GetMongoActions(),
		logAdapterOut.Log)

	accountTypeCoreAdapter = accountTypeAdapterCore.NewAdapter(storageAccountTypeOutAdapter, logAdapterOut.Log)
	accountTypeApiAdapter = accountTypeAdapterAPI.NewAdapter(
		accountTypeCoreAdapter,
		logAdapterOut.Log)

	mapAdapter = mapAdapterIn.NewAdapter(config, logAdapterOut.Log)
	postalCodeCoreAdapter = postalCodeAdapterCore.NewAdapter(logAdapterOut.Log)
	postalCodeApiAdapter = postalCodeAdapterAPI.NewAdapter(postalCodeCoreAdapter, mapAdapter, logAdapterOut.Log)

	homeschoolOnboardCoreAdapter = onboardingCoreAdapter.NewAdapter(logAdapterOut.Log, storageHomeSchoolAdapter, storageAdapter.GetMongoActions())
	homeschoolOnboardApiAdapter = onboardingApiAdapter.NewAdapter(homeschoolOnboardCoreAdapter, logAdapterOut.Log)

	httpInAdapter = httpAdapterIn.NewAdapter(
		accountTypeApiAdapter,
		postalCodeApiAdapter,
		homeschoolOnboardApiAdapter,
		logAdapterOut.Log)

	logAdapterOut.Log("info", "application starting")
	httpInAdapter.Run(logAdapterOut.HttpMiddlewareLogger)
}
