package http

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ports "github.com/codeclout/AccountEd/ports/app"
	"github.com/gofiber/fiber/v2"
)

type Adapter struct {
	api ports.AccountAPIPort
	log func(level string, msg string)
}

func NewAdapter(api ports.AccountAPIPort, logger func(level string, msg string)) *Adapter {
	return &Adapter{
		api: api,
		log: logger,
	}
}

func (a *Adapter) Run(middlewareLogger func(msg ...interface{})) {
	app := fiber.New(fiber.Config{})
	accountRoutes := a.initUserRoutes()

	app.Use(NewLoggerMiddleware(Config{
		Log: middlewareLogger,
		ShouldSkip: func(c *fiber.Ctx) bool {
			return false
		},
	}))

	app.Mount("/v1/api", accountRoutes)
	log.Fatal(app.Listen(getPort()))
}

func getPort() string {
	p, ok := os.LookupEnv("PORT")
	n, _ := strconv.Atoi(p)

	if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
		return fmt.Sprintf(":%d", n)
	}

	return ":8088"
}
