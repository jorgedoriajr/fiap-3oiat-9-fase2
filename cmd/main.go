package main

import (
	"hamburgueria/config"
	_ "hamburgueria/docs"
	"hamburgueria/internal/web/injection"
	"hamburgueria/pkg/httpclient"
	"hamburgueria/pkg/httpserver"
	"hamburgueria/pkg/sql"
	"hamburgueria/pkg/starter"
	"hamburgueria/pkg/validation"
)

// @title Hamburgueria - Grupo 9
// @version 1.0
// @description Projeto de auto atendimento para hamburgueria
func main() {
	starter.Initialize()

	var serviceConfig config.ApplicationConfig
	err := starter.UnmarshalConfig(&serviceConfig)
	if err != nil {
		panic("error on unmarshall configs")
	}
	config.SetConfig(serviceConfig)
	httpclient.Initialize()

	sql.Initialize()

	dependencyInjection := injection.NewDependencyInjection()

	server := httpserver.Builder().
		WithConfig(starter.GetHttpServerConfig()).
		WithHealthCheck(sql.GetHealthChecker()).
		WithControllers(injection.GetAllApis(dependencyInjection)...).
		WithValidator(validation.GetEchoValidator()).
		Build()

	server.Listen()
}
