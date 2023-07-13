package drivers

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	mt "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/api"
)

type Adapter struct {
	config     map[string]interface{}
	homeschool api.HomeschoolAPI
	log        *slog.Logger
}

func NewAdapter(conifg map[string]interface{}, homeschoolAPI api.HomeschoolAPI, log *slog.Logger) *Adapter {
	return &Adapter{
		config:     conifg,
		homeschool: homeschoolAPI,
		log:        log,
	}
}

// initHomeSchoolRoutes initializes the routes for the Homeschool registration process and sets specific timeouts for each route.
// It takes a *fiber.App pointer as input and returns a pointer to the updated *fiber.App.
// It checks if the SLA (Service Level Agreement) routes are defined in the config and defaults to 2000 milliseconds if not.
// The "/registration-start" route is registered with a POST method and the processRegistration function is wrapped with a timeout for handling requests.
func (a *Adapter) initHomeSchoolRoutes(app *fiber.App) *fiber.App {
	sla, ok := a.config["SLARoutes"]
	if !ok {
		a.log.Error("sla_routes not configured")
		sla = float64(2000)
	}

	b := int(sla.(float64))
	app.Post("/registration-start", timeout.NewWithContext(a.processRegistration, time.Duration(b)*time.Millisecond))

	return app
}

// InitializeAPI is a function associated with the Adapter struct.
// It takes a pointer to fiber.App as an argument and calls 'initHomeSchoolRoutes' function on 'fiber.App' pointer and returns a slice of pointers to 'fiber.App'.
// The returned slice length will depend on how many times it calls the 'initHomeSchoolRoutes' function, in this case, it's only called once.
func (a *Adapter) InitializeAPI(http *fiber.App) []*fiber.App {
	return []*fiber.App{
		a.initHomeSchoolRoutes(http),
	}
}

// processRegistration handles user registration by validating the payload and processing the pre-registration of primary members.
// It takes a *fiber.Ctx as input and returns an error if the validation or pre-registration fails, or if there's an internal server issue.
// The function makes use of sync.WaitGroup, JSON unmarshalling, and logging utility for detailed error tracking.
// Context is used to pass around values pertaining to the current request including request ID.
func (a *Adapter) processRegistration(ctx *fiber.Ctx) error {
	var in *mt.PrimaryMemberStartRegisterIn
	var wg sync.WaitGroup

	if e := mt.ValidateUsernamePayloadSize(ctx.Body()); e != nil {
		a.log.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.Status(http.StatusBadRequest))
	}

	if e := json.Unmarshal(ctx.Body(), &in); e != nil {
		a.log.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.SendStatus(http.StatusBadRequest))
	}

	if e := mt.ValidatePrimaryMember(in, &wg); e != nil {
		a.log.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.JSON(ctx.SendStatus(http.StatusBadRequest))
	}

	c := ctx.UserContext()
	cx := context.WithValue(c, mt.LogLabel("request_id"), ctx.Locals("requestid"))
	cx1 := context.WithValue(cx, mt.LogLabel("transaction_id"), sha256.Sum256([]byte(*in.Username)))

	out, x := a.HandlePreRegistration(cx1, in)
	if x != nil {
		a.log.Error(x.Error())

		return ctx.Status(http.StatusInternalServerError).JSON(errors.New("internal server error"))
	}

	return ctx.JSON(out)
}

// HandlePreRegistration takes a context and a PrimaryMemberStartRegisterIn object as input, processes the pre-registration
// of a primary member, and returns a PrimaryMemberStartRegisterOut object and an error. It leverages channels and context
// to manage request timeouts and handle errors accordingly. It also logs relevant information for debugging purposes.
func (a *Adapter) HandlePreRegistration(ctx context.Context, in *mt.PrimaryMemberStartRegisterIn) (*mt.PrimaryMemberStartRegisterOut, error) {
	ch := make(chan *mt.PrimaryMemberStartRegisterOut, 1)
	ctx, cancel := context.WithCancel(ctx)
	errorch := make(chan error, 1)

	defer cancel()

	a.homeschool.PreRegisterPrimaryMemberAPI(ctx, in, ch, errorch)

	select {
	case <-ctx.Done():
		a.log.ErrorCtx(ctx, "timeout exceeded",
			"request_id", ctx.Value(mt.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(mt.LogLabel("transaction_id"))))
		return nil, ctx.Err()

	case out := <-ch:
		a.log.Info("completed",
			"request_id", ctx.Value(mt.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(mt.LogLabel("transaction_id"))))
		return out, nil

	case e := <-errorch:
		a.log.Error(errors.Cause(e).Error(),
			"request_id", ctx.Value(mt.LogLabel("request_id")),
			"transaction_id", fmt.Sprintf("%x", ctx.Value(mt.LogLabel("transaction_id"))))
		return nil, errors.New(fiber.ErrInternalServerError.Error())
	}
}
