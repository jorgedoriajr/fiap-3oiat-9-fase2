package main

import (
	"hamburgueria/config"
	"hamburgueria/internal/application/injection"
	"hamburgueria/pkg/httpserver"
	"hamburgueria/pkg/sql"
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

	sql.Initialize()

	dependencyInjection := injection.NewDependencyInjection()

	server := httpserver.Builder().
		WithConfig(starter.GetHttpServerConfig()).
		WithHealthCheck(sql.GetHealthChecker()).
		WithControllers(injection.GetAllControllers(dependencyInjection)...).
		Build()

	server.Listen()
}
