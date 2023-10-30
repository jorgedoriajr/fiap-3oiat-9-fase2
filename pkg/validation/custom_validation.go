package validation

import "github.com/go-playground/validator"

type CustomFieldLevelValidation struct {
	Tag        string
	Validation func(fl validator.FieldLevel) bool
}
