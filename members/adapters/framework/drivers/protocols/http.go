package protocols

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type Adapter struct {
	app     *fiber.App
	config  map[string]interface{}
	monitor monitoring.Adapter
	wg      *sync.WaitGroup
}

func NewAdapter(config map[string]interface{}, app *fiber.App, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	return &Adapter{
		app:     app,
		config:  config,
		monitor: monitor,
		wg:      wg,
	}
}

func (a *Adapter) Run(port string) {
	a.monitor.LogGenericInfo("starting server")
	log.Fatal(a.app.Listen(port))
}

func (a *Adapter) StopProtocolListener(app *fiber.App) {
	a.wg.Wait()
	_ = app.Shutdown()
}
