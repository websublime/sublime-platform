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
)

func createApplication(configuration Config) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:                      configuration.Application.Prefork,
		ServerHeader:                 configuration.Application.ServerHeader,
		BodyLimit:                    configuration.Application.BodyLimit,
		CaseSensitive:                configuration.Application.CaseSensitive,
		ColorScheme:                  fiber.DefaultColors,
		CompressedFileSuffix:         configuration.Application.CompressedFileSuffix,
		Concurrency:                  configuration.Application.Concurrency,
		DisableDefaultContentType:    configuration.Application.DisableDefaultContentType,
		DisableDefaultDate:           configuration.Application.DisableDefaultDate,
		DisableHeaderNormalizing:     configuration.Application.DisableHeaderNormalizing,
		DisableKeepalive:             configuration.Application.DisableKeepalive,
		DisablePreParseMultipartForm: configuration.Application.DisablePreParseMultipartForm,
		DisableStartupMessage:        configuration.Application.DisableStartupMessage,
		ETag:                         configuration.Application.ETag,
		EnableIPValidation:           configuration.Application.EnableIPValidation,
		EnablePrintRoutes:            configuration.Application.EnablePrintRoutes,
		EnableTrustedProxyCheck:      configuration.Application.EnableTrustedProxyCheck,
		ErrorHandler:                 fiber.DefaultErrorHandler,
		GETOnly:                      configuration.Application.GETOnly,
		Immutable:                    configuration.Application.Immutable,
		ReadBufferSize:               configuration.Application.ReadBufferSize,
		StreamRequestBody:            configuration.Application.StreamRequestBody,
		StrictRouting:                configuration.Application.StrictRouting,
		UnescapePath:                 configuration.Application.UnescapePath,
		Views:                        configuration.Application.Views,
		ViewsLayout:                  configuration.Application.ViewsLayout,
		WriteBufferSize:              configuration.Application.WriteBufferSize,
	})

	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(logger.New())

	return app
}

func createServer(app *fiber.App, configuration ServerConfig) {}

func Bootstrap(configuration Config) {
	app := createApplication(configuration)

	createServer(app, configuration.Server)
}
