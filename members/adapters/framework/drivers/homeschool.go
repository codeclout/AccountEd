package drivers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/pkg/errors"

	mt "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/api"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type Adapter struct {
	config     map[string]interface{}
	homeschool api.HomeschoolAPI
	monitor    monitoring.Adapter
}

type PrimaryMemberStart = mt.PrimaryMemberStartRegisterIn
type PrimaryMemberStartOut = mt.PrimaryMemberStartRegisterOut

func NewAdapter(conifg map[string]interface{}, homeschoolAPI api.HomeschoolAPI, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:     conifg,
		homeschool: homeschoolAPI,
		monitor:    monitor,
	}
}

func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	sla, ok := a.config["SLARoutes"].(float64)
	if !ok {
		a.monitor.LogGenericError("sla_routes not configured")
		sla = float64(2000)
	}

	b := int(sla)
	app.Post("/homeschool/registration-start", timeout.NewWithContext(a.processRegistration, time.Duration(b)*time.Millisecond))

	return app
}

func (a *Adapter) InitializeAPI() []*fiber.App {
	app := fiber.New()

	return []*fiber.App{
		a.initHomeSchoolRoutes(app),
	}
}

func (a *Adapter) processRegistration(ctx *fiber.Ctx) error {
	var in *mt.PrimaryMemberStartRegisterIn
	var wg sync.WaitGroup

	if e := mt.ValidateUsernamePayloadSize(ctx.Body()); e != nil {
		a.monitor.Logger.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.Status(http.StatusBadRequest))
	}

	if e := json.Unmarshal(ctx.Body(), &in); e != nil {
		a.monitor.Logger.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.SendStatus(http.StatusBadRequest))
	}

	if e := mt.ValidatePrimaryMember(in, &wg); e != nil {
		a.monitor.Logger.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.SendStatus(http.StatusBadRequest))
	}

	c := ctx.UserContext()
	c = context.WithValue(c, a.monitor.LogLabelRequestID, ctx.Locals("requestid"))
	c = context.WithValue(c, a.monitor.LogLabelTransactionID, sha256.Sum256([]byte(*in.Username)))
	c = context.WithValue(c, a.monitor.XForwardedFor, ctx.Locals("xForwardedFor"))

	out, x := a.HandlePreRegistration(c, in)
	if x != nil {
		a.monitor.LogGenericError(x.Error())
		return ctx.JSON(ctx.Status(http.StatusInternalServerError))
	}

	return ctx.JSON(out)
}

func (a *Adapter) HandlePreRegistration(ctx context.Context, in *PrimaryMemberStart) (*PrimaryMemberStartOut, error) {
	ch := make(chan *PrimaryMemberStartOut, 1)
	ctx, cancel := context.WithCancel(ctx)
	ech := make(chan error, 1)

	defer cancel()

	a.homeschool.PreRegisterPrimaryMember(ctx, in, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogHttpError(ctx, "timeout exceeded")
		return nil, ctx.Err()

	case out := <-ch:
		a.monitor.LogHttpInfo(ctx, "completed")
		return out, nil

	case e := <-ech:
		a.monitor.LogHttpError(ctx, errors.Cause(e).Error())
		return nil, errors.Wrap(e, fiber.ErrInternalServerError.Error())
	}
}
