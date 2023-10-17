package configloader

import (
	"os"
	"regexp"

	"github.com/spf13/viper"
)

var extractorRegex = regexp.MustCompile(`^\${(?P<env>\w+):?(?P<default>.+)?}`)

func loadEnvOrDefaults(viperInstance *viper.Viper) {
	for key, value := range viperInstance.AllSettings() {
		replaceWithEnvs(viperInstance, key, value)
	}
}

func replaceWithEnvs(viperInstance *viper.Viper, key string, value interface{}) {

	if s, ok := value.(string); ok {
		env, defaultValue := getEnvAndDefault(s)
		if env != "" {
			tail := extractorRegex.ReplaceAllString(s, "")
			var configValue string
			envValue := os.Getenv(env)
			if envValue != "" {
				configValue = envValue
			} else {
				configValue = defaultValue
			}
			viperInstance.Set(key, configValue+tail)
		}

		return
	}

	if m, ok := value.(map[string]interface{}); ok {
		for currentKey, currentValue := range m {
			replaceWithEnvs(viperInstance, key+"."+currentKey, currentValue)
		}
	}
}

func getEnvAndDefault(s string) (env string, defaultValue string) {
	matches := extractorRegex.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return "", ""
	}

	groups := matches[0]
	if len(groups) > 0 {
		env = groups[1]
	}
	if len(groups) > 1 {
		defaultValue = groups[2]
	}

	return
}
