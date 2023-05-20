package drivers

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	membertypes "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type Adapter struct {
	homeschool api.HomeschoolAPI
	log        *slog.Logger
}

func NewAdapter(homeschoolAPI api.HomeschoolAPI, monitor *monitoring.Adapter) *Adapter {
	return &Adapter{
		homeschool: homeschoolAPI,
		log:        monitor.Logger,
	}
}

func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	app.Post("/registration", timeout.NewWithContext(a.processRegistration, 400*time.Millisecond))
	app.Get("/account", timeout.NewWithContext(a.processParentGuardianByAccountId, 400*time.Millisecond))

	return app
}

func (a *Adapter) processParentGuardianByAccountId(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}

func (a *Adapter) HandleParentGuardiansByAccountId(ctx context.Context, id string) *membertypes.ParentGuardianOut {
	return nil
}

func (a *Adapter) processParentGuardianById(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}

func (a *Adapter) HandleParentGuardianById(ctx context.Context, id string) *membertypes.ParentGuardianOut {
	return nil
}

func (a *Adapter) processParentGuardianByUsername(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}

func (a *Adapter) HandleParentGuardianByUsername(ctx context.Context, uname string) *membertypes.ParentGuardianOut {
	return nil
}

func (a *Adapter) processRegistration(ctx *fiber.Ctx) error {

	var in *membertypes.HomeSchoolRegisterIn
	var wg *sync.WaitGroup

	if e := membertypes.ValidatePayloadSize(ctx.Body()); e != nil {
		a.log.Error(e.Error())
		return ctx.JSON(ctx.Status(400))
	}

	if e := json.Unmarshal(ctx.Body(), &in); e != nil {
		a.log.Error(e.Error())
		return ctx.JSON(ctx.SendStatus(400))
	}

	if e := membertypes.ValidatePrimaryMember(in.PrimaryMember, wg); e != nil {
		a.log.Error(e.Error())
		return ctx.JSON(ctx.SendStatus(400))
	}

	out, x := a.HandleRegistration(ctx.UserContext(), in)
	if x != nil {
		a.log.Error(x.Error())
		ctx.Status(500)

		return ctx.JSON(errors.New("internal server error"))
	}

	return ctx.JSON(out)
}

func (a *Adapter) HandleRegistration(ctx context.Context, in *membertypes.HomeSchoolRegisterIn) (*membertypes.HomeSchoolRegisterOut, error) {
	ch := make(chan membertypes.HomeSchoolRegisterOut, 1)
	ctx, cancel := context.WithCancel(ctx)
	ech := make(chan error, 1)

	defer cancel()

	a.homeschool.RegisterAccount(ctx, in, ch, ech)

	select {
	case <-ctx.Done():
		a.log.ErrorCtx(ctx, "timeout exceeded", "")
		return nil, ctx.Err()

	case out := <-ch:
		a.log.Info("info", "")
		return &out, nil

	case e := <-ech:
		a.log.ErrorCtx(ctx, errors.Cause(e).Error())
		return nil, errors.New(fiber.ErrInternalServerError.Error())
	}
}

func (a *Adapter) processStudentByPin(ctx *fiber.Ctx) error { return errors.New("not implemented") }

func (a *Adapter) HandleStudentByPin(ctx context.Context, pin string) *membertypes.StudentOut {
	return nil
}

func (a *Adapter) processStudentByAccount(ctx *fiber.Ctx) error { return errors.New("not implemented") }

func (a *Adapter) HandleStudentByAccount(ctx context.Context, id string) *membertypes.StudentOut {
	return nil
}

func (a *Adapter) InitializeAPI(http *fiber.App) []*fiber.App {
	var out []*fiber.App

	x := a.initHomeSchoolRoutes(http)
	_ = append(out, x)

	return out
}
