package drivers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"

	"github.com/codeclout/AccountEd/members/ports/api"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

func NewHomeschoolAdapter(conifg map[string]interface{}, api api.MemberAPI, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  conifg,
		api:     api,
		monitor: monitor,
	}
}

func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	sla, ok := a.config["SLARoutes"].(float64)
	if !ok {
		a.monitor.LogGenericError("sla_routes not configured")
		sla = float64(2000)
	}

	b := int(sla)
	app.Post("/homeschool/:state/policy", timeout.NewWithContext(a.processPrimaryMemberEmail, time.Duration(b)*time.Millisecond))
	app.Post("/homeschool/:id/record", timeout.NewWithContext(a.processPrimaryMemberEmail, time.Duration(b)*time.Millisecond))
	app.Get("/homeschool/:id/stats", timeout.NewWithContext(a.processPrimaryMemberEmail, time.Duration(b)*time.Millisecond))

	return app
}

func (a *Adapter) InitializeHomeschoolAPI() []*fiber.App {
	app := fiber.New()

	return []*fiber.App{
		a.initHomeSchoolRoutes(app),
	}
}
