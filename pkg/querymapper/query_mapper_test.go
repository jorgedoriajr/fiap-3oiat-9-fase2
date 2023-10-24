package querymapper_test

import (
	"hamburgueria/pkg/querymapper"
	"testing"
	"time"
)

type customerMock struct {
	Document  string    `position:"0"`
	Name      string    `position:"1"`
	Phone     string    `position:"2"`
	Cnpj      string    `position:"3"`
	CreatedAt time.Time `position:"4"`
	UpdatedAt time.Time `position:"5"`
}

func Benchmark_getArrayFromStruct(b *testing.B) {

	source := customerMock{
		Document:  "123",
		Name:      "Test",
		Phone:     "123",
		Cnpj:      "123",
		CreatedAt: time.Now(),
	}
	for i := 0; i < b.N; i++ {
		querymapper.GetArrayOfPropertiesFrom(source)
	}
}
