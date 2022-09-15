package kernel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
)

func CreateModule(config ModuleConfig) *contracts.ModuleItem {
	app := fiber.New(fiber.Config{
		ServerHeader: config.Name,
		ErrorHandler: config.ErrorHandler,
		Views:        config.Views,
	})

	return &contracts.ModuleItem{
		Path:   config.Path,
		Module: app,
	}
}
