package drivers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/template/html/v2"

	memberT "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/members/ports/api"
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/validations"
)

type PMConfirmOut = memberT.PrimaryMemberConfirmationOut
type MErrorOut = memberT.MemberErrorOut

type cc = context.Context
type PrimaryMemberStartIn = memberT.PrimaryMemberStartRegisterIn
type PrimaryMemberStartOut = memberT.ValidatedEmailResonse

type Adapter struct {
	config  map[string]interface{}
	api     api.MemberAPI
	monitor monitoring.Adapter
}

func NewAdapter(conifg map[string]interface{}, api api.MemberAPI, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  conifg,
		api:     api,
		monitor: monitor,
	}
}

func (a *Adapter) initMemberRoutes(app *fiber.App) *fiber.App {
	sla, ok := a.config["SLARoutes"].(float64)
	if !ok {
		a.monitor.LogGenericError("sla_routes not configured")
		sla = float64(2000)
	}

	b := int(sla)

	app.Post("/register", timeout.NewWithContext(a.processPrimaryMemberEmail, time.Duration(b)*time.Millisecond))
	app.Get("/register", timeout.NewWithContext(a.processRegistrationPage, time.Duration(b)*time.Millisecond))
	app.Get("/email/confirm", timeout.NewWithContext(a.processEmailVerification, time.Duration(b)*time.Millisecond))

	return app
}

func (a *Adapter) setContextLabels(ctx *fiber.Ctx, txnValue string) context.Context {
	c := ctx.UserContext()
	c = context.WithValue(c, a.monitor.LogLabelRequestID, ctx.Locals("requestid"))
	c = context.WithValue(c, a.monitor.LogLabelTransactionID, txnValue)
	c = context.WithValue(c, a.monitor.XForwardedFor, ctx.Locals("xForwardedFor"))

	return c
}

func (a *Adapter) InitializeMemberAPI() []*fiber.App {
	workingDirectory, _ := os.Getwd()

	driver := filepath.Join(workingDirectory, "./templates")
	engine := html.New(driver, ".html")

	app := fiber.New(fiber.Config{Views: engine})

	return []*fiber.App{
		a.initMemberRoutes(app),
	}
}

func (a *Adapter) processRegistrationPage(ctx *fiber.Ctx) error {
	return ctx.Render("shell", nil)
}

func (a *Adapter) processEmailVerification(ctx *fiber.Ctx) error {
	var s string

	token := ctx.Query("t")
	if token == (s) {
		a.monitor.LogGenericError("Invalid Request: Query string required but not provided")
		return ctx.JSON(ctx.Status(fiber.StatusConflict))
	}

	c := a.setContextLabels(ctx, token)

	out, x := a.HandleEmailVerification(c, token)
	if x != nil {
		if x.Msg == "request timeout" {
			return ctx.Status(fiber.StatusGatewayTimeout).JSON(x)
		}

		return ctx.Status(http.StatusBadRequest).JSON(x)
	}

	return ctx.JSON(out)
}

func (a *Adapter) HandleEmailVerification(ctx cc, token string) (*memberT.PrimaryMemberConfirmationOut, *memberT.MemberErrorOut) {
	ch := make(chan *memberT.PrimaryMemberConfirmationOut, 1)
	ctx, cancel := context.WithCancel(ctx)
	ech := make(chan *memberT.MemberErrorOut, 1)

	defer cancel()

	a.api.GetEmailVerificationConfirmation(ctx, token, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogHttpError(ctx, "timeout exceeded")
		return nil, &MErrorOut{Error: true, Msg: "request timeout"}

	case out := <-ch:
		a.monitor.LogHttpInfo(ctx, "completed")
		return out, nil

	case e := <-ech:
		return nil, e
	}
}

func (a *Adapter) processPrimaryMemberEmail(ctx *fiber.Ctx) error {
	var in *memberT.PrimaryMemberStartRegisterIn
	var wg sync.WaitGroup

	if e := validations.ValidateUsernamePayloadSize(ctx.Body()); e != nil {
		a.monitor.Logger.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.Status(http.StatusBadRequest).JSON(memberT.MemberErrorOut{
			Error: true,
			Msg:   "invalid payload size",
		})
	}

	if e := json.Unmarshal(ctx.Body(), &in); e != nil {
		a.monitor.Logger.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.Status(http.StatusBadRequest).JSON(memberT.MemberErrorOut{
			Error: true,
			Msg:   "invalid payload",
		})
	}

	if e := validations.ValidatePrimaryMember(in, &wg); e != nil {
		a.monitor.Logger.Error(e.Error(), "request_id", ctx.Locals("requestid"))
		return ctx.Status(http.StatusBadRequest).JSON(memberT.MemberErrorOut{
			Error: true,
			Msg:   "invalid primary member address",
		})
	}

	hash := sha256.Sum256([]byte(*in.Username))
	str := hex.EncodeToString(hash[:])
	c := a.setContextLabels(ctx, str)

	out, x := a.HandlePrimaryMemberValidity(c, in)
	if x != nil {
		if x.Msg == "request timeout" {
			return ctx.Status(fiber.StatusGatewayTimeout).JSON(x)
		}

		return ctx.Status(http.StatusBadRequest).JSON(x)
	}

	return ctx.JSON(out)
}

func (a *Adapter) HandlePrimaryMemberValidity(ctx cc, in *PrimaryMemberStartIn) (*PrimaryMemberStartOut, *MErrorOut) {
	ch := make(chan *PrimaryMemberStartOut, 1)
	ctx, cancel := context.WithCancel(ctx)
	ech := make(chan *memberT.MemberErrorOut, 1)

	defer cancel()

	a.api.VerifyPrimaryMemberEmail(ctx, in, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogHttpError(ctx, "timeout exceeded")
		return nil, &memberT.MemberErrorOut{Error: true, Msg: "request timeout"}

	case out := <-ch:
		a.monitor.LogHttpInfo(ctx, "completed")
		return out, nil

	case e := <-ech:
		return nil, e
	}
}
