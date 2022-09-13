package kernel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Application() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:                      false,
		ServerHeader:                 "ws-platform",
		BodyLimit:                    4 * 1024 * 1024,
		CaseSensitive:                false,
		ColorScheme:                  fiber.DefaultColors,
		CompressedFileSuffix:         ".gz",
		Concurrency:                  256 * 1024,
		DisableDefaultContentType:    false,
		DisableDefaultDate:           false,
		DisableHeaderNormalizing:     false,
		DisableKeepalive:             false,
		DisablePreParseMultipartForm: false,
		DisableStartupMessage:        false,
		ETag:                         false,
		EnableIPValidation:           false,
		EnablePrintRoutes:            false,
		EnableTrustedProxyCheck:      false,
		ErrorHandler:                 fiber.DefaultErrorHandler,
		GETOnly:                      false,
		//IdleTimeout: nil,
		Immutable:         false,
		ReadBufferSize:    4096,
		StreamRequestBody: false,
		StrictRouting:     false,
		//TrustedProxies: []string*__*,
		UnescapePath:    false,
		Views:           nil,
		ViewsLayout:     "",
		WriteBufferSize: 4096,
	})

	app.Use(recover.New())
	app.Use(pprof.New())
	app.Use(logger.New())

	return app
}
