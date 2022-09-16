package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/kernel"
)

type HomeController struct {
	Foundation *kernel.Foundation
}

func HelloController(c *fiber.Ctx) error {
	return c.Render("modules/website/views/home", fiber.Map{})
}
