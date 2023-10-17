package configloader

import "fmt"

type missingRequiredConfigError string

func (str missingRequiredConfigError) Error() string {
	return fmt.Sprintf("missing required configuration: %q", string(str))
}

func validate(o options) error {

	if len(o.configPaths) == 0 {
		return missingRequiredConfigError("ConfigPath")
	}

	if o.configType == "" {
		return missingRequiredConfigError("ConfigType")
	}

	if len(o.configName) == 0 {
		return missingRequiredConfigError("ConfigName")
	}

	return nil
}
