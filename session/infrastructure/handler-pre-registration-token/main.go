package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	apiAdapter "handler-pre-registration-token/adapters/api"
	coreAdapter "handler-pre-registration-token/adapters/core"
	drivenAdapter "handler-pre-registration-token/adapters/framework/driven"
	"handler-pre-registration-token/adapters/framework/drivers"
)

func main() {
	monitor := monitoring.NewAdapter()

	driven := drivenAdapter.NewAdapter(*monitor)
	core := coreAdapter.NewAdapter(*monitor)

	api := apiAdapter.NewAdapter(core, driven, *monitor)
	driver := drivers.NewAdapter(api, *monitor)

	lambda.Start(driver.GenerateMemberToken)
}
