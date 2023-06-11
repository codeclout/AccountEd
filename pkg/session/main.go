package main

import (
	"github.com/codeclout/AccountEd/pkg/monitoring"
	configuration "github.com/codeclout/AccountEd/pkg/session/adapter/framework/drivers/server"
)

func main() {
	monitor := monitoring.NewAdapter()
	go monitor.Initialize()

	sessionConfiguration := configuration.NewAdapter(monitor.Logger)
	internalConfig := sessionConfiguration.LoadSessionConfig()

}
