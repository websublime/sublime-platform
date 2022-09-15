package pkg

import (
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/foundation/kernel"
	"github.com/websublime/platform/modules/website"
)

func Setup(config kernel.Config) {
	registerModule(website.NewModuleWebsite())
}

func registerModule(modules ...contracts.ModuleInterface) {
	for _, module := range modules {
		module.RegisterModule()
	}
}
