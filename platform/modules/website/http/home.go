package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/kernel"
)

type HomeController struct {
	AppContext *kernel.Foundation
}

func HelloController(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
