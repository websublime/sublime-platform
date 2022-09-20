package contracts

type Context struct {
	Modules     Modules
	Services    Services
	Config      *Config
	Environment EnvironmentConfig
}
