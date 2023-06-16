package main

import (
	"github.com/codeclout/AccountEd/pkg/monitoring"
	configuration "github.com/codeclout/AccountEd/pkg/session/adapter/framework/drivers/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	monitor := monitoring.NewAdapter()
	wg.Add(1)
	go monitor.Initialize(&wg)

	sessionConfiguration := configuration.NewAdapter(monitor.Logger)
	internalConfig := sessionConfiguration.LoadSessionConfig()

}
