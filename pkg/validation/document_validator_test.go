package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCPF(t *testing.T) {
	testCases := []struct {
		cpf      string
		expected bool
	}{
		{cpf: "529.982.247-25", expected: true},
		{cpf: "123.456.789-09", expected: true},
		{cpf: "77416899249", expected: true},
		{cpf: "123.456.789-10", expected: false},
		{cpf: "11122233344", expected: false},
		{cpf: "05251270976", expected: true},
	}

	for _, tc := range testCases {
		result := ValidateCPF(tc.cpf)
		assert.Equal(t, tc.expected, result)
	}
}
