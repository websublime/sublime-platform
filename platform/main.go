package main

import (
	"github.com/websublime/foundation"
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/platform/modules/website"
)

func main() {
	config := contracts.Config{
		Application: contracts.ApplicationConfig{
			ServerHeader: "ws-platform",
		},
		Modules: contracts.Modules{
			website.RegisterWebsiteModule(),
		},
	}

	_, bootserver := foundation.Start(&config, false)

	bootserver(config.Server)
}
