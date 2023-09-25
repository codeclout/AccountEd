package main

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/codeclout/AccountEd/notifications/infrastructure/adapters/framework/drivers/runner"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

func main() {
	app := cdktf.NewApp(nil)

	monitor := monitoring.NewAdapter()
	infrastructure := runner.NewAdapter(app, *monitor)

	infrastructure.InitializeInfrastructure()
	app.Synth()
}
