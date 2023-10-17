package configloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	t.Run("Validate ConfigPath", func(t *testing.T) {

		expectedErr := missingRequiredConfigError("ConfigPath")
		o := options{
			configType: "yaml",
			configName: []string{"config"},
		}

		err := validate(o)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("Validate ConfigType", func(t *testing.T) {

		expectedErr := missingRequiredConfigError("ConfigType")
		o := options{
			configPaths: []string{"./config"},
			configName:  []string{"config"},
		}

		err := validate(o)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("Validate ConfigName", func(t *testing.T) {

		expectedErr := missingRequiredConfigError("ConfigName")
		o := options{
			configPaths: []string{"./config"},
			configType:  "yaml",
		}

		err := validate(o)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("No error", func(t *testing.T) {

		o := options{
			configPaths: []string{"./config"},
			configType:  "yaml",
			configName:  []string{"config"},
		}

		err := validate(o)
		assert.NoError(t, err)
	})
}
