/*
Copyright © 2022 Websublime.dev organization@websublime.dev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package kernel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/websublime/foundation/contracts"
)

type Foundation struct {
	Modules     contracts.Modules
	Services    contracts.Services
	Config      *Config
	Environment EnvironmentConfig
}

func createApp(configuration ApplicationConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:                      configuration.Prefork,
		ServerHeader:                 configuration.ServerHeader,
		BodyLimit:                    configuration.BodyLimit,
		CaseSensitive:                configuration.CaseSensitive,
		ColorScheme:                  fiber.DefaultColors,
		CompressedFileSuffix:         configuration.CompressedFileSuffix,
		Concurrency:                  configuration.Concurrency,
		DisableDefaultContentType:    configuration.DisableDefaultContentType,
		DisableDefaultDate:           configuration.DisableDefaultDate,
		DisableHeaderNormalizing:     configuration.DisableHeaderNormalizing,
		DisableKeepalive:             configuration.DisableKeepalive,
		DisablePreParseMultipartForm: configuration.DisablePreParseMultipartForm,
		DisableStartupMessage:        configuration.DisableStartupMessage,
		ETag:                         configuration.ETag,
		EnableIPValidation:           configuration.EnableIPValidation,
		EnablePrintRoutes:            configuration.EnablePrintRoutes,
		EnableTrustedProxyCheck:      configuration.EnableTrustedProxyCheck,
		ErrorHandler:                 fiber.DefaultErrorHandler,
		GETOnly:                      configuration.GETOnly,
		Immutable:                    configuration.Immutable,
		ReadBufferSize:               configuration.ReadBufferSize,
		StreamRequestBody:            configuration.StreamRequestBody,
		StrictRouting:                configuration.StrictRouting,
		UnescapePath:                 configuration.UnescapePath,
		Views:                        configuration.Views,
		ViewsLayout:                  configuration.ViewsLayout,
		WriteBufferSize:              configuration.WriteBufferSize,
	})

	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(logger.New())

	return app
}

func NewApplication(configuration *Config) (*Foundation, *fiber.App) {
	env := GetEnvironmentConfig()

	configuration.Application.Prefork = env.IsProduction

	if configuration.Server.Host == "" {
		configuration.Server.Host = env.WsHost
	}

	if configuration.Server.Port == "" {
		configuration.Server.Port = env.WsPort
	}

	foundation := &Foundation{
		Config:      configuration,
		Environment: env,
		Modules:     GetModules(),
		Services:    GetServices(),
	}

	app := createApp(configuration.Application)

	for _, moduleItem := range foundation.Modules {
		if moduleItem.Config.Active {
			app.Mount(moduleItem.Path, moduleItem.ModuleApp)
		}
	}

	return foundation, app
}
