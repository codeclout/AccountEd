package protocols

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type ProtocolPort interface {
	Initialize(api []*fiber.App) (*fiber.App, string)
	PostInit(app *fiber.App, wg *sync.WaitGroup)
	StopProtocolListener(app *fiber.App)
}
