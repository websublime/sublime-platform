package contracts

import (
	"github.com/gofiber/fiber/v2"
)

type ModuleConfig struct {
	Path         string             `json:"path"`
	Name         string             `json:"name"`
	ErrorHandler fiber.ErrorHandler `json:"-"`
	Views        fiber.Views        `json:"-"`
	Active       bool               `json:"active"`
	ViewsLayout  string
}

type ModuleItem struct {
	Path      string
	ModuleApp *fiber.App
	Config    ModuleConfig
}

type Module struct {
	Item        *ModuleItem
	Group       fiber.Router
	Controllers Controllers
}

type Modules []*Module
