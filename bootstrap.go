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
package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/websublime/sublime-platform/api"
	"github.com/websublime/sublime-platform/config"
	"github.com/websublime/sublime-platform/pkg"
	"github.com/websublime/sublime-platform/utils"
)

type Router interface {
	InstallRouter(app *fiber.App)
}

func bootstrap(env *config.Environment) *fiber.App {
	engine := html.NewFileSystem(http.FS(pkg.ViewFiles), ".html")

	app := fiber.New(fiber.Config{
		ServerHeader: "ws-platform",
		Prefork:      env.IsProduction,
		Views:        engine,
		ErrorHandler: errorHandler,
	})

	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(logger.New())

	return app
}

func installRouter(app *fiber.App) {
	setup(app, api.NewApi(), pkg.NewWebsite(), pkg.NewConsole())
}

func setup(app *fiber.App, router ...Router) {
	for _, group := range router {
		group.InstallRouter(app)
	}
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError
	message := "Internal server error"
	exception := utils.ErrorServerUnknown

	// Retreive the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	if e, ok := err.(*utils.Exception); ok {
		code = e.Code
		message = e.Message
		exception = e.Exception
	}

	// Send custom error page
	ctx.Set(fiber.HeaderContentType, "application/json")

	return ctx.Status(code).JSON(fiber.Map{
		"status": code,
		"error":  message,
		"code":   exception,
	})
}
