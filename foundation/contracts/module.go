package contracts

import "github.com/gofiber/fiber/v2"

type ModuleItem struct {
	Path   string
	Module *fiber.App
}

type Modules []*ModuleItem
