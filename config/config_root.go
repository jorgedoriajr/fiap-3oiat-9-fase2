package config

type Root struct {
	Application Application
	HttpServer  HttpServerConfig
	Databases   map[string]DatabaseConfig
}
