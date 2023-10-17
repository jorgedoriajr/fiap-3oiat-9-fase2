package configloader

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Config represents a loaded config
type Config interface {
	// Unmarshal unmarshals config into given struct
	Unmarshal(v interface{}) error
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	UnmarshalKey(key string, output interface{}) error
	// WriteConfigAs writes current configuration to a given filename.
	WriteConfigAs(filename string) error
	// Set sets the value for the key in the override register.
	SetKeyValue(key string, value interface{})
}

type config struct {
	viperInstance *viper.Viper
	opts          options
}

func getViper(opts options) (*viper.Viper, error) {
	viperInstance := viper.New()
	viperInstance.SetConfigName(opts.configName[0])
	viperInstance.SetConfigType(opts.configType)
	for i := range opts.configPaths {
		viperInstance.AddConfigPath(opts.configPaths[i])
	}
	for _, configName := range opts.configName {
		viperInstance.SetConfigName(configName)
		err := viperInstance.MergeInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return nil, err
			}
		}
	}

	for i := range opts.configReaders {
		err := viperInstance.MergeConfig(opts.configReaders[i])
		if err != nil {
			return nil, err
		}
	}

	if opts.automaticEnv {
		viperInstance.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
		viperInstance.AutomaticEnv()
	}

	return viperInstance, nil
}

// Load loads config files
func Load(options ...opt) (Config, error) {
	opts, err := getOptions(options...)
	if err != nil {
		return nil, err
	}

	viperInstance, err := getViper(opts)
	if err != nil {
		return nil, err
	}

	loadEnvOrDefaults(viperInstance)

	return &config{
		viperInstance: viperInstance,
		opts:          opts,
	}, nil
}

func (c *config) Unmarshal(output interface{}) error {
	if c.opts.defaultFromStruct {
		c.defaultFromStruct("", output)
	}
	return c.viperInstance.Unmarshal(&output)
}

func (c *config) UnmarshalKey(key string, output interface{}) error {
	if c.opts.defaultFromStruct {
		c.defaultFromStruct(key, output)
	}
	return c.viperInstance.UnmarshalKey(key, &output)
}

func (c *config) defaultFromStruct(baseKey string, value interface{}) {
	cfgPtr := reflect.ValueOf(value)
	cfgElem := cfgPtr.Elem()
	cfgElemType := cfgElem.Type()

	for i := 0; i < cfgElem.NumField(); i++ {
		f := cfgElemType.Field(i)
		fValue := cfgElem.Field(i)
		if key, ok := f.Tag.Lookup("mapstructure"); ok {
			currentKey := baseKey + key
			if f.Type.Kind() == reflect.Struct {
				c.defaultFromStruct(currentKey+".", fValue.Addr().Interface())
			}
			if defaultValue, ok := f.Tag.Lookup("default"); ok {
				c.viperInstance.SetDefault(currentKey, defaultValue)
			}
		}
	}
}

func (c *config) WriteConfigAs(filename string) error {
	return c.viperInstance.WriteConfigAs(filename)
}

func (c *config) SetKeyValue(key string, value interface{}) {
	c.viperInstance.Set(key, value)
}
