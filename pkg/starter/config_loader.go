package starter

import (
	"fmt"
	"hamburgueria/internal/util/configloader"
	"hamburgueria/internal/util/slice"
	"os"
)

type configOptions struct {
	ConfigPath string
	ConfigType string
	ConfigName []string
	ProfileEnv string
}

func defaultConfigOptions() configOptions {
	return configOptions{
		ConfigPath: "./config",
		ConfigType: "yaml",
		ConfigName: []string{"config"},
		ProfileEnv: "GO_PROFILE",
	}
}

/*
SetConfigPath overrides default config path value

Usage:

	application.Initialize(
	  application.SetConfigPath("path/to/config"), // without config file name
	)

default: ./config
*/
func SetConfigPath(path string) opt {
	return func(o *options) error {
		o.configOptions.ConfigPath = path
		return nil
	}
}

/*
SetConfigType overrides default config type value

Usage:

	application.Initialize(
	  application.SetConfigType("json"),
	)

default: yaml
*/
func SetConfigType(configType string) opt {
	return func(o *options) error {
		o.configOptions.ConfigType = configType
		return nil
	}
}

/*
SetConfigName overrides default config name value

Usage:

	application.Initialize(
	  application.SetConfigName("config", "config-prod"), // without extension
	)

default: ["config"]
*/
func SetConfigName(configName ...string) opt {
	return func(o *options) error {
		o.configOptions.ConfigName = configName
		return nil
	}
}

/*
SetProfileEnv overrides default profile environment value (read from os environment).
If set, application will try to load config file for given profile.

Example:

If config name is set to "config" and profile env evaluates to "dev", application will load config file and append config-dev to configuration source.

Usage:

	application.Initialize(
	  application.SetProfileEnv("ENV_PROFILE"),
	)

default: GO_PROFILE
*/
func SetProfileEnv(profileEnv string) opt {
	return func(o *options) error {
		o.configOptions.ProfileEnv = profileEnv
		return nil
	}
}

func initializeConfig() {
	c, err := loadConfig(appInstance.configOptions)
	if err != nil {
		panic(fmt.Sprintf("error loading config: %q", err))
	}

	appInstance.config = c
	err = c.Unmarshal(&appInstance.configRoot)
	if err != nil {
		panic(fmt.Sprintf("error unmarshal config: %q", err))
	}
}

func loadConfig(opt configOptions) (configloader.Config, error) {

	if opt.ProfileEnv != "" {
		if len(opt.ConfigName) > 0 {

			profileConfig := fmt.Sprintf("%s-%s", opt.ConfigName[0], os.Getenv(opt.ProfileEnv))
			if !slice.Contains(opt.ConfigName, profileConfig) {
				opt.ConfigName = append(opt.ConfigName, profileConfig)
			}
		}
	}

	return configloader.Load(
		configloader.WithConfigType(opt.ConfigType),
		configloader.WithConfigPath(opt.ConfigPath),
		configloader.WithConfigName(opt.ConfigName[1:]...), // Prevent duplication of "config" in list.
	)
}
