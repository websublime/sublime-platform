package main

import (
	"github.com/websublime/foundation"
	"github.com/websublime/foundation/kernel"
	"github.com/websublime/platform/modules/website"
)

func main() {
	website.NewModuleWebsite().InstallModule()

	bootserver := foundation.Start(kernel.Config{
		Application: kernel.ApplicationConfig{
			ServerHeader: "ws-platform",
		},
		Server: kernel.ServerConfig{
			Host: "localhost",
			Port: "8080",
			Tls:  false,
		},
	}, false)

	bootserver()
}
