package foundation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/foundation/kernel"
)

func createServer(app *fiber.App, configuration kernel.ServerConfig) {
	host := fmt.Sprintf("%s:%s", configuration.Host, configuration.Port)

	if configuration.Tls {
		app.ListenTLS(host, configuration.TlsCertFile, configuration.TlsKeyFile)
	} else {
		app.Listen(host)
	}
}

func Start(config *kernel.Config, bootServer bool) (*kernel.Foundation, func(configuration kernel.ServerConfig)) {
	foundation, app := kernel.NewApplication(config)

	if bootServer {
		createServer(app, config.Server)
	}

	return foundation, func(configuration kernel.ServerConfig) {
		createServer(app, configuration)
	}
}

func Setup(modules ...contracts.ModuleInterface) {
	for _, module := range modules {
		module.RegisterModule()
	}
}
