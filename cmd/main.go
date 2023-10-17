package main

import (
	"hamburgueria/config"
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

	server := httpserver.Builder().
		WithConfig(starter.GetHttpServerConfig()).
		Build()

	server.Listen()
}
