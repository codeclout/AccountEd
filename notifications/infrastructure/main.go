package main

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"cdk.tf/go/stack/adapters/framework/drivers/runner"
	"github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

func main() {
	app := cdktf.NewApp(nil)

	monitor := drivers.NewAdapter()
	infrastructure := runner.NewAdapter(app, monitor.Logger)

	infrastructure.InitializeInfrastructure()
	app.Synth()
}
