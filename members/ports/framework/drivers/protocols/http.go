package protocols

import "github.com/gofiber/fiber/v2"

type MemberProtocolHTTPPort interface {
	Run(port string)
	StopProtocolListener(app *fiber.App)
}
