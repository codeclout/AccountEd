package drivers

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	mt "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type Adapter struct {
	config     map[string]interface{}
	homeschool api.HomeschoolAPI
	log        *slog.Logger
}

func NewAdapter(homeschoolAPI api.HomeschoolAPI, monitor *monitoring.Adapter, runtimeConfig map[string]interface{}) *Adapter {
	return &Adapter{
		config:     runtimeConfig,
		homeschool: homeschoolAPI,
		log:        monitor.Logger,
	}
}

func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	b, _ := strconv.Atoi(a.config["sla_routePerformance"].(string))
	app.Post("/registration-start", timeout.NewWithContext(a.processRegistration, time.Duration(b)*time.Millisecond))

	return app
}

func (a *Adapter) InitializeAPI(http *fiber.App) []*fiber.App {
	var out []*fiber.App

	x := a.initHomeSchoolRoutes(http)
	_ = append(out, x)

	return out
}

func (a *Adapter) processRegistration(ctx *fiber.Ctx) error {
	var in *mt.PrimaryMemberStartRegisterIn
	var wg *sync.WaitGroup

	if e := mt.ValidatePayloadSize(ctx.Body()); e != nil {
		a.log.Error(e.Error())
		return ctx.JSON(ctx.Status(400))
	}

	if e := json.Unmarshal(ctx.Body(), &in); e != nil {
		a.log.Error(e.Error())
		return ctx.JSON(ctx.SendStatus(400))
	}

	if e := mt.ValidatePrimaryMember(in, wg); e != nil {
		a.log.Error(e.Error())
		return ctx.JSON(ctx.SendStatus(400))
	}

	out, x := a.HandlePreRegistration(ctx.UserContext(), in)
	if x != nil {
		a.log.Error(x.Error())
		ctx.Status(500)

		return ctx.JSON(errors.New("internal server error"))
	}

	return ctx.JSON(out)
}

func (a *Adapter) HandlePreRegistration(ctx context.Context, in *mt.PrimaryMemberStartRegisterIn) (*mt.PrimaryMemberStartRegisterOut, error) {
	ch := make(chan *mt.PrimaryMemberStartRegisterOut, 1)
	ctx, cancel := context.WithCancel(ctx)
	errorch := make(chan error, 1)

	defer cancel()

	a.homeschool.PreRegisterPrimaryMember(ctx, in, ch, errorch)

	select {
	case <-ctx.Done():
		a.log.ErrorCtx(ctx, "timeout exceeded", "") // @TODO - transaction ID
		return nil, ctx.Err()

	case out := <-ch:
		a.log.Info("info", "log the transaction ID") // @TODO - transaction ID
		return out, nil

	case e := <-errorch:
		a.log.ErrorCtx(ctx, errors.Cause(e).Error()) // @TODO - transaction ID
		return nil, errors.New(fiber.ErrInternalServerError.Error())
	}
}
