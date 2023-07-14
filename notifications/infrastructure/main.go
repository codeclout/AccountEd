package main

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"cdk.tf/go/stack/adapters/framework/drivers/runner"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

func main() {
	app := cdktf.NewApp(nil)

	monitor := monitoring.NewAdapter()
	infrastructure := runner.NewAdapter(app, monitor.Logger)

	infrastructure.InitializeInfrastructure()
	app.Synth()
}
