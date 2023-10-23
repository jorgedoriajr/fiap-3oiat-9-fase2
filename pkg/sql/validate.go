package sql

import (
	"fmt"
)

type missingRequiredConfigError string

func (str missingRequiredConfigError) Error() string {
	return fmt.Sprintf("missing required configuration: %q", string(str))
}

func validateConfig(conf Config) error {
	if conf.Host == "" {
		return missingRequiredConfigError("database.host")
	}
	if conf.DatabaseName == "" {
		return missingRequiredConfigError("database.databaseName")
	}
	return nil
}
