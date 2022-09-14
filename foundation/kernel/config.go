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

import "github.com/gofiber/fiber/v2"

type EnvironmentConfig struct {
	WsEnvironment string `mapstructure:"WS_ENV"`
	WsApiUrl      string `mapstructure:"WS_API_URL"`
	WsApiKey      string `mapstructure:"WS_API_KEY"`
	WsApiSecret   string `mapstructure:"WS_API_SECRET"`
	WsHost        string `mapstructure:"WS_HOST"`
	WsPort        string `mapstructure:"WS_PORT"`
	IsProduction  bool
}

type ApplicationConfig struct {
	Prefork                      bool               `json:"prefork"`
	ServerHeader                 string             `json:"serverHeader"`
	BodyLimit                    int                `json:"bodyLimit"`
	CaseSensitive                bool               `json:"caseSensitive"`
	CompressedFileSuffix         string             `json:"compressFileSuffix"`
	Concurrency                  int                `json:"concurrency"`
	DisableDefaultContentType    bool               `json:"disableDefaultContentType"`
	DisableDefaultDate           bool               `json:"disableDefaultDate"`
	DisableHeaderNormalizing     bool               `json:"disableHeaderNormalize"`
	DisableKeepalive             bool               `json:"disableKeepAlive"`
	DisablePreParseMultipartForm bool               `json:"disablePreParseMultipartForm"`
	DisableStartupMessage        bool               `json:"disableStartupMessage"`
	ETag                         bool               `json:"eTag"`
	EnableIPValidation           bool               `json:"enableIPValidation"`
	EnablePrintRoutes            bool               `json:"enablePrintRoutes"`
	EnableTrustedProxyCheck      bool               `json:"enableTrustedProxyCheck"`
	ErrorHandler                 fiber.ErrorHandler `json:"-"`
	GETOnly                      bool               `json:"getOnly"`
	Immutable                    bool               `json:"immutable"`
	ReadBufferSize               int                `json:"readBufferSize"`
	StreamRequestBody            bool               `json:"streamRequestBody"`
	StrictRouting                bool               `json:"strictRouting"`
	UnescapePath                 bool               `json:"unescapePath"`
	Views                        fiber.Views        `json:"-"`
	ViewsLayout                  string             `json:"viewsLayout"`
	WriteBufferSize              int                `json:"writeBufferSize"`
}

type ServerConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Tls         bool   `json:"tls"`
	TlsCertFile string `json:"tlsCertFile"`
	TlsKeyFile  string `json:"tlsKeyFile"`
}

type ModuleConfig struct {
	ServerHeader string             `json:"serverHeader"`
	ErrorHandler fiber.ErrorHandler `json:"-"`
	Views        fiber.Views        `json:"-"`
}

type Config struct {
	Environment EnvironmentConfig `json:"environment"`
	Application ApplicationConfig `json:"application"`
	Server      ServerConfig      `json:"server"`
}
