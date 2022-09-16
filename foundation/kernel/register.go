package kernel

import "github.com/websublime/foundation/contracts"

var (
	modules  = contracts.Modules{}
	services = contracts.Services{}
)

func GetModules() contracts.Modules {
	return modules
}

func AddModule(module *contracts.Module) {
	modules = append(modules, module)
}

func GetServices() contracts.Services {
	return services
}

func AddService(service *contracts.ServiceItem) {
	services = append(services, service)
}
