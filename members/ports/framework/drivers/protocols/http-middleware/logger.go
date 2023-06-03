package httpmiddleware

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type Config struct {
	Log        func(msg ...interface{})
	ShouldSkip func(c *fiber.Ctx) bool
}

var Internals = Config{
	Log:        nil,
	ShouldSkip: nil,
}

func NewLoggerMiddleware(c ...Config) fiber.Handler {
	var r string
	config := setConfig(c...)

	return func(ctx *fiber.Ctx) error {
		req := ctx.Request()
		res := ctx.Response()

		if config.ShouldSkip(ctx) {
			return ctx.Next()
		}

		uuid, e := exec.Command("uuidgen").Output()
		if e != nil {
			log.Fatal(e)
		}

		r = string(req.Header.Peek(fiber.HeaderXRequestID))
		if r == "" {
			r = strings.Trim(string(uuid), "\n")
		}

		x := string(req.Header.Peek(fiber.HeaderXForwardedFor))
		if x == "" {
			x = ctx.IP()
		}

		config.Log("incoming request", slog.Group("request",
			slog.String("host", string(req.Host())),
			slog.String("method", ctx.Method()),
			slog.String("path", string(req.URI().Path())),
			slog.String("request_id", r),
			slog.Int("status", res.StatusCode()),
			slog.String("uri", string(req.RequestURI())),
			slog.String("x_forward_for", x),
		))

		return ctx.Next()
	}
}

func setConfig(c ...Config) Config {
	if len(c) == 0 {
		return Internals
	}

	config := c[0]

	if config.Log == nil {
		config.ShouldSkip = func(c *fiber.Ctx) bool {
			return true
		}
	}

	return config
}
