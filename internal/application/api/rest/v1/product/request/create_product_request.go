package request

import (
	"hamburgueria/internal/modules/product/usecase/command"
)

type CreateProductRequest struct {
	Name        string              `json:"name" validator:"required"`
	Description string              `json:"description" validator:"required"`
	Category    string              `json:"category" validator:"required"`
	Menu        bool                `json:"menu" validator:"required"`
	ImgPath     string              `json:"imgPath" validator:"required"`
	Ingredients []IngredientRequest `json:"ingredients" validator:"required"`
}

type IngredientRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name" validator:"required"`
	Quantity int    `json:"quantity" validator:"required"`
	Amount   int    `json:"amount" validator:"required"`
	Type     string `json:"type" validator:"required,ingredientType"`
}

func (cp CreateProductRequest) ToCommand() command.CreateProductCommand {
	return *command.NewCreateProductCommand(
		cp.Name, cp.Description, cp.Category, cp.Menu, toIngredients(cp.Ingredients), cp.ImgPath,
	)
}

func toIngredients(ingredients []IngredientRequest) []command.Ingredient {
	var ingredientsCmd []command.Ingredient
	for _, ingredient := range ingredients {
		ingredientsCmd = append(ingredientsCmd, command.Ingredient{
			ID:       ingredient.ID,
			Quantity: ingredient.Quantity,
			//Amount:   ingredient.Amount,
			//Type:     command.GetIngredientTypeByName(ingredient.Type),
		})
	}
	return ingredientsCmd
}
