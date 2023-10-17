package starter

import (
	"hamburgueria/config"
	"hamburgueria/internal/util/configloader"
	"sync"
)

var (
	appInstance   *app
	opts          *options
	initOnce      sync.Once
	isInitialized = false
)

// Represents the application options.
type options struct {
	configOptions *configOptions
}

type app struct {
	configOptions configOptions
	config        configloader.Config
	configRoot    *config.Root
}

type opt func(*options) error

func Initialize(options ...opt) {
	ensureNotInitialized()
	ensureCreated()
	applyOptions(options...)
	initializeConfig()

	isInitialized = true
}

func UnmarshalConfig(v any) error {
	ensureInitialized()
	return appInstance.config.Unmarshal(&v)
}

func applyOptions(options ...opt) {
	opts.configOptions = &appInstance.configOptions
	for _, op := range options {
		err := op(opts)
		if err != nil {
			panic(err)
		}
	}
}
