package validation

import (
	"github.com/go-playground/validator"
)

type DefaultValidator struct {
	Validator *validator.Validate
}

func (d *DefaultValidator) Validate(i interface{}) error {
	return d.Validator.Struct(i)
}
