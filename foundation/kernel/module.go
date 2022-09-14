package kernel

import "github.com/gofiber/fiber/v2"

func CreateModule(config ModuleConfig) *fiber.App {
	return fiber.New(fiber.Config{
		ServerHeader: config.ServerHeader,
		ErrorHandler: config.ErrorHandler,
		Views:        config.Views,
	})
}
