package main

import (
	"hamburgueria/config"
	"hamburgueria/internal/application/injection"
	"hamburgueria/internal/application/rest"
	"hamburgueria/pkg/httpserver"
	"hamburgueria/pkg/starter"
)

func main() {

	starter.Initialize()

	var serviceConfig config.ApplicationConfig
	err := starter.UnmarshalConfig(&serviceConfig)
	if err != nil {
		panic("error on unmarshall configs")
	}
	config.SetConfig(serviceConfig)

	dependencyInjection := injection.NewDependencyInjection(serviceConfig)

	server := httpserver.Builder().
		WithConfig(starter.GetHttpServerConfig()).
		WithControllers(rest.GetAllControllers(dependencyInjection)).
		Build()

	server.Listen()
}
