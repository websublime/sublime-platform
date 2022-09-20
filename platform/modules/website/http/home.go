package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
)

type Home contracts.Base

func HelloController(c *fiber.Ctx) error {
	return c.Render("modules/website/views/home", fiber.Map{
		"Url": "localhost",
	})
}

func (home *Home) Hello(c *fiber.Ctx) error {
	return c.Render("modules/website/views/home", fiber.Map{
		"Url": "localhost",
	})
}
