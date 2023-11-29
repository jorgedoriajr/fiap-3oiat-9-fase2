package config

// DatabaseConfig represents sql database configuration
/*
LogLevel
Silent silent log level - 1
Error error log level - 2
Warn warn log level - 3
Info info log level - 4
*/
type DatabaseConfig struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
	MaxPoolSize  int
	Schema       string
	LogLevel     int
}
