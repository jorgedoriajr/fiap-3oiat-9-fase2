package config

import (
	"hamburgueria/pkg/httpclient"
)

type Root struct {
	Application Application
	HttpServer  HttpServerConfig
	Databases   map[string]DatabaseConfig
	HttpClients map[string]httpclient.Config
}
