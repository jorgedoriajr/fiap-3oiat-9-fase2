package slice_test

import (
	"hamburgueria/internal/util/slice"
	"testing"

	"github.com/stretchr/testify/assert"
)

type KeyValue struct {
	Key, Value string
}

func TestToMap(t *testing.T) {

	t.Run("ToMap with primitive input", func(t *testing.T) {
		input := []string{"value", "another value"}

		m := slice.ToMap(
			input,
			func(s string) string { return s },
			func(s string) int { return len(s) },
		)

		assert.Equal(t, 2, len(m))
		assert.Equal(t, len("value"), m["value"])
		assert.Equal(t, len("another value"), m["another value"])
	})

	t.Run("ToMap with struct input", func(t *testing.T) {
		input := []KeyValue{
			{
				Key:   "Key1",
				Value: "Value 1",
			},
			{
				Key:   "Key2",
				Value: "Value 2",
			},
		}

		m := slice.ToMap(
			input,
			func(kv KeyValue) string { return kv.Key },
			func(kv KeyValue) string { return kv.Value },
		)

		assert.Equal(t, 2, len(m))
		assert.Equal(t, "Value 1", m["Key1"])
		assert.Equal(t, "Value 2", m["Key2"])
	})

}
