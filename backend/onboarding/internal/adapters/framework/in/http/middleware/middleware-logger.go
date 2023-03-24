package middleware

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

	return func(c *fiber.Ctx) error {
		req := c.Request()
		res := c.Response()

		if config.ShouldSkip(c) == true {
			return c.Next()
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
			x = c.IP()
		}

		config.Log("host", string(req.Host()),
			"method", c.Method(),
			"path", string(req.URI().Path()),
			"requestId", r,
			"status", strconv.Itoa(res.StatusCode()),
			"uri", string(req.RequestURI()),
			"xforwarded", x)

		return c.Next()
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
