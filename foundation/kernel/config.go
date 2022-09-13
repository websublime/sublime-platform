package kernel

type Environment struct {
	WsEnvironment string `mapstructure:"WS_ENV"`
	WsApiUrl      string `mapstructure:"WS_API_URL"`
	WsApiKey      string `mapstructure:"WS_API_KEY"`
	WsApiSecret   string `mapstructure:"WS_API_SECRET"`
	WsHost        string `mapstructure:"WS_HOST"`
	WsPort        string `mapstructure:"WS_PORT"`
	IsProduction  bool
}

type Config struct {
	Prefork bool `json:"prefork"`
	ServerHeader string `json:"serverHeader"`
	BodyLimit int `json:"bodyLimit"`
	CaseSensitive bool `json:"caseSensitive"`
	CompressedFileSuffix string `json:"compressFileSuffix"`
	Concurrency int `json:"concurrency"`
	DisableDefaultContentType bool `json:"disableDefaultContentType"`
	DisableDefaultDate bool `json:"disableDefaultDate"`
	DisableHeaderNormalizing bool `json:"disableHeaderNormalize"`
	DisableKeepalive bool `json:"disableKeepAlive"`
	DisablePreParseMultipartForm bool `json:"disablePreParseMultipartForm"`
	DisableStartupMessage bool `json:"disableStartupMessage"`
	ETag bool `json:"eTag"`
	EnableIPValidation: false,
	EnablePrintRoutes: false,
	EnableTrustedProxyCheck: false,
	ErrorHandler: fiber.DefaultErrorHandler,
	GETOnly: false,
	//IdleTimeout: nil,
	Immutable: false,
	ReadBufferSize: 4096,
	StreamRequestBody: false,
	StrictRouting: false,
	//TrustedProxies: []string*__*,
	UnescapePath: false,
	Views: nil,
	ViewsLayout: "",
	WriteBufferSize: 4096,
}