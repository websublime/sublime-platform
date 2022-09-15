package main

import (
	"github.com/websublime/foundation"
	"github.com/websublime/foundation/kernel"
	"github.com/websublime/platform/pkg"
)

func main() {
	config := kernel.Config{
		Application: kernel.ApplicationConfig{
			ServerHeader: "ws-platform",
		},
	}

	pkg.Setup(config)

	_, bootserver := foundation.Start(&config, false)

	bootserver(config.Server)
}
