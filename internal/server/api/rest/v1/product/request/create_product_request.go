package request

import (
	"hamburgueria/internal/modules/product/usecase/command"
)

type ProductRequest struct {
	Name        string              `json:"name" validator:"required"`
	Description string              `json:"description" validator:"required"`
	Category    string              `json:"category" validator:"required"`
	Menu        bool                `json:"menu" validator:"required"`
	ImgPath     string              `json:"imgPath" validator:"required"`
	Ingredients []IngredientRequest `json:"ingredients" validator:"required"`
}

type IngredientRequest struct {
	Number   int `json:"number"`
	Quantity int `json:"quantity" validator:"required"`
}

func (cp ProductRequest) ToCommand() command.CreateProductCommand {
	return *command.NewCreateProductCommand(
		cp.Name, cp.Description, cp.Category, cp.Menu, toIngredients(cp.Ingredients), cp.ImgPath,
	)
}

func toIngredients(ingredients []IngredientRequest) []command.Ingredient {
	var ingredientsCmd []command.Ingredient
	for _, ingredient := range ingredients {
		ingredientsCmd = append(ingredientsCmd, command.Ingredient{
			Number:   ingredient.Number,
			Quantity: ingredient.Quantity,
		})
	}
	return ingredientsCmd
}
