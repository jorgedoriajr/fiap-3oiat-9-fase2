package httpclient

import (
	"fmt"
	"hamburgueria/pkg/starter"
	"strings"
	"sync"
)

const defaultClientName = "_default"

var (
	clients        map[string]Client
	initializeOnce sync.Once
)

func Initialize() {
	initializeOnce.Do(func() {
		clients = make(map[string]Client)
		conf := starter.GetHttpClientsConfig()

		var defaultConf Config
		if c, exists := conf[defaultClientName]; exists {
			defaultConf = c
		} else {
			defaultConf = Config{}
		}

		for k, v := range conf {
			if k == defaultClientName {
				continue
			}

			mergedConfig := mergeConfig(defaultConf, v)

			c, err := NewClient(k, mergedConfig)
			if err != nil {
				panic(err)
			}
			clients[strings.ToLower(k)] = c
		}
	})
}

func GetClient(name string) Client {
	if c, ok := clients[strings.ToLower(name)]; ok {
		return c
	}

	panic(fmt.Sprintf("no http client found for %q", name))
}
