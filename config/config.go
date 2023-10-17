package config

type ApplicationConfig struct {
}

var configInstance ApplicationConfig

func SetConfig(c ApplicationConfig) {
	configInstance = c
}

func GetConfig() ApplicationConfig {
	return configInstance
}
