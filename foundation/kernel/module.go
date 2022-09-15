package kernel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
)

func CreateModule(config contracts.ModuleConfig) contracts.Module {
	app := fiber.New(fiber.Config{
		ServerHeader: config.Name,
		ErrorHandler: config.ErrorHandler,
		Views:        config.Views,
	})

	return contracts.Module{
		Item: &contracts.ModuleItem{
			Path:      config.Path,
			ModuleApp: app,
			Config:    config,
		},
	}
}
