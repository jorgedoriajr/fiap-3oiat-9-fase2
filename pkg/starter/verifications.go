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
		panic("application already initialized")
	}
}

func ensureInitialized() {
	if !isInitialized {
		panic("application must be initialized")
	}
}
