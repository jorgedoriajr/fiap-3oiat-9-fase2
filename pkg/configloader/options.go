package configloader

import "io"

type options struct {
	configPaths       []string
	configType        string
	configName        []string
	defaultFromStruct bool
	automaticEnv      bool
	configReaders     []io.Reader
}

type opt func(*options) error

/*
WithConfigName sets config names

default: ["config"]
*/
func WithConfigName(configNames ...string) opt {
	return func(o *options) error {
		o.configName = append(o.configName, configNames...)
		return nil
	}
}

/*
WithConfigPath sets config path

default: ./config
*/
func WithConfigPath(path string) opt {
	return func(o *options) error {
		o.configPaths = append(o.configPaths, path)
		return nil
	}
}

/*
WithConfigPath sets config path

default: ./config
*/
func WithConfigPaths(paths ...string) opt {
	return func(o *options) error {
		o.configPaths = append(o.configPaths, paths...)
		return nil
	}
}

/*
WithConfigType sets config type

default: ./yaml
*/
func WithConfigType(configType string) opt {
	return func(o *options) error {
		o.configType = configType
		return nil
	}
}

/*
WithDefaultFromStruct reads values default from struct, also enables
automatic env

default: false
*/
func WithDefaultFromStruct() opt {
	return func(o *options) error {
		o.defaultFromStruct = true
		o.automaticEnv = true
		return nil
	}
}

/*
WithAutomaticEnv enables values from environment variables without the
need to specify their keys in the config file

default: false
*/
func WithAutomaticEnv() opt {
	return func(o *options) error {
		o.automaticEnv = true
		return nil
	}
}

/*
WithConfigReaders sets config readers. Config readers are
useful as they are binary representations of the configs, meaning
that, when using them, we no longer need to have a physical file
in our docker image.

default: nil
*/
func WithConfigReaders(readers ...io.Reader) opt {
	return func(o *options) error {
		o.configReaders = readers
		return nil
	}
}

func defaultOptions() options {
	return options{
		configPaths: []string{"./config"},
		configType:  "yaml",
		configName:  []string{"config"},
	}
}

func getOptions(opts ...opt) (options, error) {
	o := defaultOptions()
	for _, op := range opts {
		err := op(&o)
		if err != nil {
			return options{}, err
		}
	}

	err := validate(o)
	if err != nil {
		return options{}, err
	}

	return o, nil
}
