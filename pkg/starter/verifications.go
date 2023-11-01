package starter

import "hamburgueria/config"

// Creates server with default configuration.
func ensureCreated() {
	initOnce.Do(func() {
		opts = &options{}
		appInstance = &app{
			configOptions: defaultConfigOptions(),
			configRoot:    &config.Root{},
		}
	})
}

func ensureNotInitialized() {
	if isInitialized {
		panic("server already initialized")
	}
}

func ensureInitialized() {
	if !isInitialized {
		panic("server must be initialized")
	}
}
