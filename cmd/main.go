package main

import (
	"hamburgueria/config"
	_ "hamburgueria/docs"
	"hamburgueria/internal/application/injection"
	"hamburgueria/pkg/httpserver"
	"hamburgueria/pkg/sql"
	"hamburgueria/pkg/starter"
	"hamburgueria/pkg/validation"
)

// @title Hamburgueria - Grupo 9
// @version 1.0
// @description Projeto de auto atendimento para hamburgueria
// @host https://localhost/8080
// @BasePath /v1
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
		WithValidator(validation.GetEchoValidator()).
		Build()

	server.Listen()
}
