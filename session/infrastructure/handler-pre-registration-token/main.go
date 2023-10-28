package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	apiAdapter "handler-pre-registration-token/adapters/api"
	coreAdapter "handler-pre-registration-token/adapters/core"
	drivenAdapter "handler-pre-registration-token/adapters/framework/driven"
	"handler-pre-registration-token/adapters/framework/drivers"
	"handler-pre-registration-token/configuration"
)

func main() {
	monitor := monitoring.NewAdapter()
	configurationAdapter := configuration.NewAdapter("./config.hcl", *monitor)
	config := configurationAdapter.LoadStorageConfig()

	driven := drivenAdapter.NewAdapter(*monitor)
	core := coreAdapter.NewAdapter(*monitor)

	api := apiAdapter.NewAdapter(core, driven, *monitor)
	driver := drivers.NewAdapter(api, *config, *monitor)

	lambda.Start(driver.GenerateMemberToken)
}
