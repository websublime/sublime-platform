package website

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/foundation/kernel"
	"github.com/websublime/platform/modules/website/http"
)

type WebsiteModule struct {
	Module      *contracts.Module
	Group       fiber.Router
	Controllers contracts.Controllers
}

func NewModuleWebsite() *WebsiteModule {
	mod := kernel.CreateModule(contracts.ModuleConfig{
		Path:   "/",
		Name:   "website",
		Active: true,
	})

	kernel.AddModule(mod.Item)

	return &WebsiteModule{
		Module: &mod,
		Controllers: contracts.Controllers{
			&contracts.ControllerItem{
				Provider: http.HelloController,
				Path:     "",
				Method:   "GET",
				Name:     "home",
			},
		},
	}
}

func (mod WebsiteModule) RegisterModule() {
	mod.Group = mod.Module.Item.ModuleApp.Group("").Name(fmt.Sprintf("%s-group", mod.Module.Item.Config.Name))

	for _, controller := range mod.Controllers {
		mod.Group.Add(controller.Method, controller.Path, controller.Provider).Name(controller.Name)
	}
}
