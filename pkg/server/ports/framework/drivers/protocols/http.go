package protocols

import "github.com/gofiber/fiber/v2"

type ProtocolHTTP interface {
  Initialize(api []*fiber.App)
  PostInit(app *fiber.App)
  StopProtocolListener(app *fiber.App)
}
