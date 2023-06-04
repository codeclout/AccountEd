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

	mt "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/api"
)

type label string

type Adapter struct {
	config     map[string]interface{}
	homeschool api.HomeschoolAPI
	log        *slog.Logger
}

func NewAdapter(homeschoolAPI api.HomeschoolAPI, log *slog.Logger, runtimeConfig map[string]interface{}) *Adapter {
	return &Adapter{
		config:     runtimeConfig,
		homeschool: homeschoolAPI,
		log:        log,
	}
}

func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	b := int(a.config["sla_routes"].(float64))
	app.Post("/registration-start", timeout.NewWithContext(a.processRegistration, time.Duration(b)*time.Millisecond))

	return app
}

func (a *Adapter) InitializeAPI(http *fiber.App) []*fiber.App {
	var out []*fiber.App

	x := a.initHomeSchoolRoutes(http)
	out = append(out, x)

	return out
}

func (a *Adapter) processRegistration(ctx *fiber.Ctx) error {
	var in *mt.PrimaryMemberStartRegisterIn
	var wg sync.WaitGroup

	if e := mt.ValidateUsernamePayloadSize(ctx.Body()); e != nil {
		a.log.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.Status(400))
	}

	if e := json.Unmarshal(ctx.Body(), &in); e != nil {
		a.log.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.SendStatus(400))
	}

	if e := mt.ValidatePrimaryMember(in, &wg); e != nil {
		a.log.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.SendStatus(400))
	}

	c := ctx.UserContext()
	cstr := label("requestid")
	cx := context.WithValue(c, cstr, ctx.Locals("requestid"))

	out, x := a.HandlePreRegistration(cx, in)
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
		a.log.ErrorCtx(ctx, "timeout exceeded", "request_id", ctx.Value(label("requestid")))
		return nil, ctx.Err()

	case out := <-ch:
		a.log.Info("completed", "request_id", ctx.Value(label("requestid")))
		return out, nil

	case e := <-errorch:
		a.log.Error(errors.Cause(e).Error(), "request_id", ctx.Value(label("requestid")))
		return nil, errors.New(fiber.ErrInternalServerError.Error())
	}
}
