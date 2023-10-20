package configloader

import (
	_ "embed"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed test_files/config.yaml
	testConfig string
	//go:embed test_files/config-profile.yaml
	profileConfig string
)

type SimpleConfig struct {
	StringValue string
	IntValue    int
	InnerConfig SimpleInnerConfig
}

type AnotherConfig struct {
	InnerConfig   SimpleInnerConfig
	AnotherString string
}

type SimpleInnerConfig struct {
	InnerStringValue string
}

type ConfigWithTaggedDefault struct {
	Value       string                 `mapstructure:"value" default:"abc"`
	InnerConfig InnerWithTaggedDefault `mapstructure:"inner-config"`
}

type InnerWithTaggedDefault struct {
	Value        string                      `mapstructure:"value" default:"cde"`
	InnerConfig  InnerInnerWithTaggedDefault `mapstructure:"inner-config"`
	AnotherValue string                      `mapstructure:"another-value" default:"efg"`
}

type InnerInnerWithTaggedDefault struct {
	Value        string `mapstructure:"value" default:"hij"`
	AnotherValue string `mapstructure:"another-value" default:"klm"`
}

func TestUnmarshal(t *testing.T) {

	t.Run("Config load", func(t *testing.T) {
		var simpleConfig SimpleConfig

		c, err := Load(
			WithConfigPath("./test_files"),
			WithConfigType("yaml"),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "string1", simpleConfig.StringValue)
		assert.Equal(t, 2, simpleConfig.IntValue)
		assert.Equal(t, "string2", simpleConfig.InnerConfig.InnerStringValue)

		var another AnotherConfig
		err = c.Unmarshal(&another)

		assert.NoError(t, err)
		assert.NotNil(t, another)
		assert.Equal(t, "string2", another.InnerConfig.InnerStringValue)
		assert.Equal(t, "another string", another.AnotherString)
	})

	t.Run("Config load with profile overriding", func(t *testing.T) {
		var simpleConfig SimpleConfig

		c, err := Load(
			WithConfigName("config", "config-profile"),
			WithConfigPath("./test_files"),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "specificString", simpleConfig.StringValue)
		assert.Equal(t, 2, simpleConfig.IntValue)
		assert.Equal(t, "string2", simpleConfig.InnerConfig.InnerStringValue)
	})

	t.Run("Config load with env and defaults", func(t *testing.T) {

		os.Setenv("STRING_ENV_VAR2", "string value 2")
		os.Setenv("INT_ENV_VAR", "15")

		var simpleConfig SimpleConfig

		c, err := Load(
			WithConfigName("config-env"),
			WithConfigName("config-env-profile"),
			WithConfigPath("./test_files"),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "string value 2", simpleConfig.StringValue)
		assert.Equal(t, 15, simpleConfig.IntValue)
		assert.Equal(t, "defaultValue plus the rest of string", simpleConfig.InnerConfig.InnerStringValue)
	})

	t.Run("Config load unknown with tagged defaults with option WithDefaultFromStruct", func(t *testing.T) {

		var simpleConfig ConfigWithTaggedDefault

		c, err := Load(
			WithConfigName("config-env"),
			WithConfigName("config-env-profile"),
			WithConfigPath("./test_files"),
			WithDefaultFromStruct(),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "abc", simpleConfig.Value)
		assert.Equal(t, "cde", simpleConfig.InnerConfig.Value)
		assert.Equal(t, "efg", simpleConfig.InnerConfig.AnotherValue)
		assert.Equal(t, "hij", simpleConfig.InnerConfig.InnerConfig.Value)
		assert.Equal(t, "klm", simpleConfig.InnerConfig.InnerConfig.AnotherValue)
	})

	t.Run("Config load with tagged defaults with option WithDefaultFromStruct", func(t *testing.T) {

		var simpleConfig ConfigWithTaggedDefault

		os.Setenv("ENV_ANOTHER_INNER_CONFIG_VALUE", "envvar")
		os.Setenv("INNER_CONFIG_INNER_CONFIG_ANOTHER_VALUE", "morePriority")
		os.Setenv("PRIORITY", "lessPriority")
		os.Setenv("INNER_CONFIG_INNER_CONFIG_VALUE", "fromEnvAuto")

		c, err := Load(
			WithConfigName("config-values"),
			WithConfigName("config-env-profile"),
			WithConfigPath("./test_files"),
			WithDefaultFromStruct(),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "value", simpleConfig.Value)
		assert.Equal(t, "value_str", simpleConfig.InnerConfig.Value)
		assert.Equal(t, "envvar", simpleConfig.InnerConfig.AnotherValue)
		assert.Equal(t, "fromEnvAuto", simpleConfig.InnerConfig.InnerConfig.Value)
		assert.Equal(t, "morePriority", simpleConfig.InnerConfig.InnerConfig.AnotherValue)
	})

	t.Run("Config load with automatic env vars with option WithAutomaticEnv", func(t *testing.T) {
		var simpleConfig SimpleConfig

		os.Setenv("STRINGVALUE", "env-var")
		os.Setenv("INTVALUE", "99")
		os.Setenv("INNERCONFIG_INNERSTRINGVALUE", "another-env-var")

		c, err := Load(
			WithConfigName("config"),
			WithConfigPath("./test_files"),
			WithAutomaticEnv(),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "env-var", simpleConfig.StringValue)
		assert.Equal(t, 99, simpleConfig.IntValue)
		assert.Equal(t, "another-env-var", simpleConfig.InnerConfig.InnerStringValue)
	})

	t.Run("Config load with config readers", func(t *testing.T) {
		var simpleConfig SimpleConfig

		c, err := Load(
			WithConfigReaders(strings.NewReader(testConfig), strings.NewReader(profileConfig)),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "specificString", simpleConfig.StringValue)
		assert.Equal(t, 2, simpleConfig.IntValue)
		assert.Equal(t, "string2", simpleConfig.InnerConfig.InnerStringValue)
	})
	t.Run("Config load with config readers and regular config files and defaults", func(t *testing.T) {
		var simpleConfig SimpleConfig

		c, err := Load(
			WithConfigPath("./test_files"),
			WithConfigType("yaml"),
			WithConfigReaders(strings.NewReader(testConfig), strings.NewReader(profileConfig)),
		)
		assert.NoError(t, err)

		err = c.Unmarshal(&simpleConfig)

		assert.NoError(t, err)
		assert.NotNil(t, simpleConfig)
		assert.Equal(t, "specificString", simpleConfig.StringValue)
		assert.Equal(t, 2, simpleConfig.IntValue)
		assert.Equal(t, "string2", simpleConfig.InnerConfig.InnerStringValue)
	})
}

func TestUnmarshalKey(t *testing.T) {
	t.Run("UnmarshalKey should unmarshal the given key", func(t *testing.T) {
		var inner SimpleInnerConfig

		c, err := Load(
			WithConfigType("yaml"),
		)
		assert.NoError(t, err)

		err = c.UnmarshalKey("InnerConfig", &inner)

		assert.NoError(t, err)
		assert.NotNil(t, inner)
		assert.Equal(t, "string2", inner.InnerStringValue)
	})
}
