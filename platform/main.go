package main

import (
	"github.com/websublime/foundation"
	"github.com/websublime/foundation/kernel"
	"github.com/websublime/platform/modules/website"
)

func main() {
	config := kernel.Config{
		Application: kernel.ApplicationConfig{
			ServerHeader: "ws-platform",
		},
	}

	foundation.Setup(website.NewModuleWebsite())

	_, bootserver := foundation.Start(&config, false)

	bootserver(config.Server)
}
