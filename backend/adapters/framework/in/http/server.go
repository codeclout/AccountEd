package http

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codeclout/AccountEd/adapters/framework/in/http/middleware"
	port "github.com/codeclout/AccountEd/ports/api/account-types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

type l func(level, msg string)
type p port.UserAccountTypeApiPort

type Adapter struct {
	api p
	log l
}

func NewAdapter(api p, logger l) *Adapter {
	return &Adapter{
		api: api,
		log: logger,
	}
}

func (a *Adapter) Run(middlewareLogger func(msg ...interface{})) {
	app := fiber.New(fiber.Config{})
	accountRoutes := a.initUserRoutes()

	app.Use(middleware.NewLoggerMiddleware(middleware.Config{
		Log: middlewareLogger,
		ShouldSkip: func(c *fiber.Ctx) bool {
			return false
		},
	}))
	app.Use(etag.New())
	//app.Use(AirCollision412())

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
