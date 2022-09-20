package kernel

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
)

func CreateModule(config contracts.ModuleConfig, base contracts.Base) contracts.Module {
	app := fiber.New(fiber.Config{
		ServerHeader: config.Name,
		ErrorHandler: config.ErrorHandler,
		Views:        config.Views,
		ViewsLayout:  config.ViewsLayout,
	})

	module := contracts.Module{
		Base: &base,
		Item: &contracts.ModuleItem{
			Path:     config.Path,
			Instance: app,
			Config:   config,
		},
	}

	AddModule(&module)

	return module
}

func RegisterModule(config contracts.ModuleConfig, base contracts.Base) *contracts.Module {
	mod := CreateModule(contracts.ModuleConfig{
		Path:        config.Path,
		Name:        config.Name,
		Active:      config.Active,
		Views:       config.Views,
		ViewsLayout: config.ViewsLayout,
	}, base)

	mod.Base.Group = mod.Item.Instance.Group("").Name(fmt.Sprintf("%s-group", strings.ToLower(mod.Item.Config.Name)))

	for _, controller := range mod.Base.Controllers {
		mod.Base.Group.Add(
			controller.Item.Config.Method,
			controller.Item.Config.Path,
			controller.Item.Config.Handler,
		).Name(controller.Item.Config.Name)
	}

	/*
		mod.Group = mod.Item.ModuleApp.Group("").Name(fmt.Sprintf("%s-group", strings.ToLower(mod.Item.Config.Name)))
		var controls contracts.Controllers

		for _, cfg := range controllers {
			control := CreateControler(contracts.ControllerConfig{
				Provider: cfg.Provider,
				Path:     cfg.Path,
				Method:   cfg.Method,
				Name:     cfg.Name,
			})

			controls = append(controls, &control)
		}

		mod.Controllers = controls

		for _, controller := range mod.Controllers {
			mod.Group.Add(
				controller.Item.Config.Method, controller.Item.Config.Path, controller.Item.Config.Provider,
			).Name(controller.Item.Config.Name)
		}
	*/

	return &mod
}
