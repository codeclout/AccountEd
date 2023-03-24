package http

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	postalCodePortAPI "github.com/codeclout/AccountEd/internal/ports/api/postal-codes"
	"github.com/codeclout/AccountEd/onboarding/internal/adapters/framework/in/http/middleware"
	accountTypeApiPort "github.com/codeclout/AccountEd/onboarding/internal/ports/api/account-types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

type logger func(level, msg string)

type Adapter struct {
	postalCodeApi      postalCodePortAPI.PostalCodeApiPort
	userAccountTypeApi accountTypeApiPort.UserAccountTypeApiPort
	log                logger
}

func NewAdapter(accountTypeApi accountTypeApiPort.UserAccountTypeApiPort, postalCodeApi postalCodePortAPI.PostalCodeApiPort, logger logger) *Adapter {
	return &Adapter{
		log:                logger,
		postalCodeApi:      postalCodeApi,
		userAccountTypeApi: accountTypeApi,
	}
}

func (a *Adapter) Run(middlewareLogger func(msg ...interface{})) {
	app := fiber.New(fiber.Config{})
	api := fiber.New()

	_ = a.initUserAccountTypeRoutes(api)
	// _ = a.initPostalCodeRoutes(api)

	app.Use(middleware.NewLoggerMiddleware(middleware.Config{
		Log: middlewareLogger,
		ShouldSkip: func(c *fiber.Ctx) bool {
			return false
		},
	}))
	app.Use(etag.New())
	//app.Use(AirCollision412())

	app.Mount("/v1/api", api)

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
