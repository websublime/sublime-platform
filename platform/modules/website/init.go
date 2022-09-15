package website

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/foundation/kernel"
)

type WebsiteModule struct {
	Item *contracts.ModuleItem
}

func NewModuleWebsite() *WebsiteModule {
	mod := kernel.CreateModule(kernel.ModuleConfig{
		Path: "/",
		Name: "Platform-Website",
	})

	kernel.AddModule(mod)

	return &WebsiteModule{
		Item: mod,
	}
}

func (mod WebsiteModule) InstallModule() {
	mod.Item.Module.Group("", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
}
