package contracts

import "github.com/gofiber/fiber/v2"

type ControllerItem struct {
	Path     string
	Name     string
	Provider fiber.Handler
	Method   string
}

type Controllers []*ControllerItem
