package http_middleware

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
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
	config := setConfig(c...)

	return func(ctx *fiber.Ctx) error {
		req := ctx.Request()
		res := ctx.Response()

		if config.ShouldSkip(ctx) == true {
			return ctx.Next()
		}

		uuid, e := exec.Command("uuidgen").Output()
		if e != nil {
			log.Fatal(e)
		}

		r := string(req.Header.Peek(fiber.HeaderXRequestID))
		if r == "" {
			r = strings.Trim(string(uuid), "\n")
		}

		x := string(req.Header.Peek(fiber.HeaderXForwardedFor))
		if x == "" {
			x = ctx.IP()
		}

		config.Log("host", string(req.Host()),
			"method", ctx.Method(),
			"path", string(req.URI().Path()),
			"requestId", r,
			"status", strconv.Itoa(res.StatusCode()),
			"uri", string(req.RequestURI()),
			"xforwarded", x)

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
