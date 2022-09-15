package foundation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

func Start(config kernel.Config, bootServer bool) func() {
	app := kernel.CreateApplication(config)

	if bootServer {
		createServer(app, config.Server)
	} else {
		return func() {
			createServer(app, config.Server)
		}
	}

	return nil
}
