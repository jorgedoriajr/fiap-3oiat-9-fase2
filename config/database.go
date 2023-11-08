package config

// DatabaseConfig represents sql database configuration
// QueryMode should be passed with value simple_protocol if you are considering use pgBouncer.
type DatabaseConfig struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
	MaxPoolSize  int
	Schema       string
	QueryMode    string
}
