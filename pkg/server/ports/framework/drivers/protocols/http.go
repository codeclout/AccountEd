package protocols

import "github.com/gofiber/fiber/v2"

type ProtocolPort interface {
  Initialize(api []*fiber.App) *fiber.App
  PostInit(app *fiber.App)
  StopProtocolListener(app *fiber.App)
}
