package starter

import "hamburgueria/config"

// GetAppConfig return the loaded application config
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
