package starter

import (
	"hamburgueria/config"
)

// GetAppConfig return the loaded server config
func GetAppConfig() config.Application {
	ensureInitialized()
	return appInstance.configRoot.Application
}

// GetHttpServerConfig return the http server config
func GetHttpServerConfig() config.HttpServerConfig {
	ensureInitialized()
	return appInstance.configRoot.HttpServer
}

// GetDatabasesConfig return the loaded database config
func GetDatabasesConfig() map[string]config.DatabaseConfig {
	ensureInitialized()
	return appInstance.configRoot.Databases
}

// GetHttpClientsConfig return the loaded http client config
func GetHttpClientsConfig() map[string]config.HttpClientConfig {
	ensureInitialized()
	return appInstance.configRoot.HttpClients
}
