package validation

import (
	"github.com/go-playground/validator"
	"hamburgueria/internal/modules/product/domain/valueobject"
)

func ValidateProductCategory(category string) bool {
	allowedCategories := []valueobject.ProductCategory{
		valueobject.Dish,
		valueobject.Drink,
		valueobject.Dessert,
	}

	inputType := valueobject.ProductCategory(category)

	for _, t := range allowedCategories {
		if t == inputType {
			return true
		}
	}
	return false
}

func EchoValidateProductCategory(fl validator.FieldLevel) bool {
	inputType := valueobject.ProductCategory(fl.Field().String())
	return ValidateProductCategory(string(inputType))
}
