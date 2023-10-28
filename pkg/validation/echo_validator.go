package validation

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EchoValidator struct {
	Validator *validator.Validate
}

func (cv *EchoValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (cv *EchoValidator) RegisterValidators(customFieldValidations []CustomFieldLevelValidation) {
	for _, custom := range customFieldValidations {
		err := cv.Validator.RegisterValidation(custom.Tag, custom.Validation)
		if err != nil {
			panic(err.Error())
		}
	}

}

func GetEchoValidator() *EchoValidator {
	vld := &EchoValidator{Validator: validator.New()}
	vld.RegisterValidators([]CustomFieldLevelValidation{
		{"ingredientType", ValidateIngredientType},
		{"productCategory", EchoValidateProductCategory},
		{"cpf", ValidateCPFType},
	})
	return vld
}
