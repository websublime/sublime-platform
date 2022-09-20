package kernel

import "github.com/websublime/foundation/contracts"

func CreateControler(config contracts.ControllerConfig) *contracts.Controller {
	return &contracts.Controller{
		Item: contracts.ControllerItem{
			Config: contracts.ControllerConfig{
				Path:    config.Path,
				Name:    config.Name,
				Handler: config.Handler,
				Method:  config.Method,
			},
		},
	}
}
