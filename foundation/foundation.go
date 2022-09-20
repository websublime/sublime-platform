package foundation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/foundation/kernel"
)

func createServer(app *fiber.App, configuration contracts.ServerConfig) {
	host := fmt.Sprintf("%s:%s", configuration.Host, configuration.Port)

	if configuration.Tls {
		app.ListenTLS(host, configuration.TlsCertFile, configuration.TlsKeyFile)
	} else {
		app.Listen(host)
	}
}

func bootModules(ctx *contracts.Context, modules contracts.Modules) {
	for _, module := range modules {
		module.Base.DefineContext(ctx)
	}
}

func Start(config *contracts.Config, bootServer bool) (*contracts.Context, func(configuration contracts.ServerConfig)) {
	foundation, app := kernel.NewApplication(config)

	bootModules(foundation, config.Modules)

	if bootServer {
		createServer(app, config.Server)
	}

	return foundation, func(configuration contracts.ServerConfig) {
		createServer(app, configuration)
	}
}
