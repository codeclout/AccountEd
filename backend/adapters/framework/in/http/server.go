package http

import (
	"log"

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
	app := fiber.New()

	app.Post("/", a.CreateAccountType)
	log.Fatal(app.Listen(":3000"))
}
