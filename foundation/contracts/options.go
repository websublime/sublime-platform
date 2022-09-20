package contracts

import "github.com/gofiber/fiber/v2"

type EnvironmentConfig struct {
	Environment  string `mapstructure:"WS_ENV"`
	ApiUrl       string `mapstructure:"WS_API_URL"`
	ApiKey       string `mapstructure:"WS_API_KEY"`
	ApiSecret    string `mapstructure:"WS_API_SECRET"`
	Host         string `mapstructure:"WS_HOST"`
	Port         string `mapstructure:"WS_PORT"`
	Url          string `mapstructure:"WS_URL"`
	IsProduction bool
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

type Config struct {
	Application ApplicationConfig `json:"application"`
	Server      ServerConfig      `json:"server"`
	Modules     Modules
}
