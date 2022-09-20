package contracts

import "github.com/gofiber/fiber/v2"

type ControllerConfig struct {
	Path    string
	Name    string
	Handler fiber.Handler
	Method  string
}

type ControllerItem struct {
	Config ControllerConfig
}

type Controllers []*Controller

type Controller struct {
	Item ControllerItem
}

type Base struct {
	Group       fiber.Router
	Controllers Controllers
	Foundation  *Context
}

func (base *Base) DefineContext(context *Context) {
	base.Foundation = context
}
