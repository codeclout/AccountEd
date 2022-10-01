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
}

func NewAdapter(api ports.AccountAPIPort) *Adapter {
	return &Adapter{api: api}
}

func (a Adapter) Run() {
	app := fiber.New(fiber.Config{})

	app.Post("/", a.CreateAccountType)
	log.Fatal(app.Listen(GetPort()))
}

func GetPort() string {
	p, ok := os.LookupEnv("PORT")
	n, _ := strconv.Atoi(p)

	if ok && len(strings.TrimSpace(p)) >= 4 && n >= 1024 && n <= 65535 {
		return fmt.Sprintf(":%d", n)
	}

	return ":8088"
}
