package validation

import (
	"github.com/go-playground/validator"
	"hamburgueria/internal/modules/product/usecase/command"
)

func ValidateIngredientType(fl validator.FieldLevel) bool {
	allowedTypes := []command.IngredientType{
		command.Protein,
		command.VegetableAndSalad,
		command.Sauces,
		command.Cheeses,
	}

	inputType := command.IngredientType(fl.Field().String())

	for _, t := range allowedTypes {
		if t == inputType {
			return true
		}
	}

	return false
}
