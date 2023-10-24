package querymapper_test

import (
	"github.com/stretchr/testify/assert"
	"hamburgueria/pkg/querymapper"
	"testing"
	"time"
)

type customerMock struct {
	Document  string    `position:"0"`
	Name      string    `position:"1"`
	Phone     string    `position:"2"`
	CreatedAt time.Time `position:"4"`
	UpdatedAt time.Time `position:"5"`
}

func TestOrderingCheck(t *testing.T) {
	t.Run("should return properties in correct order", func(t *testing.T) {

		source := customerMock{
			Phone:     "21222222",
			Name:      "Test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Document:  "123",
		}

		result := querymapper.GetArrayOfPropertiesFrom(source)

		assert.Equal(t, "Test", result[1])
		assert.Equal(t, "123", result[0])
		assert.Equal(t, "21222222", result[2])
	})
}

func Benchmark_getArrayFromStruct(b *testing.B) {

	source := customerMock{
		Document:  "123",
		Name:      "Test",
		Phone:     "123",
		CreatedAt: time.Now(),
	}
	for i := 0; i < b.N; i++ {
		querymapper.GetArrayOfPropertiesFrom(source)
	}
}
