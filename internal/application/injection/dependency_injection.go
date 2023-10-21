package injection

import (
	"hamburgueria/config"
	"hamburgueria/internal/application/rest"
)

type DependencyInjection struct {
	CustomerController *rest.CustomerController
}

func NewDependencyInjection(config config.ApplicationConfig) DependencyInjection {
	return DependencyInjection{}
}
