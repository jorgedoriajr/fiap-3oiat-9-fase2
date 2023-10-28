package request

import (
	"hamburgueria/internal/modules/product/domain/valueobject"
	"hamburgueria/internal/modules/product/usecase/command"
)

type CreateProductRequest struct {
	Name        string                      `json:"name" validator:"required"`
	Description string                      `json:"description" validator:"required"`
	Category    valueobject.ProductCategory `json:"category" validator:"required,productCategory"`
	Menu        bool                        `json:"menu" validator:"required"`
	Ingredients []IngredientRequest         `json:"ingredients" validator:"required"`
}

type IngredientRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name" validator:"required"`
	Quantity int    `json:"quantity" validator:"required"`
	Amount   int    `json:"amount" validator:"required"`
	Type     string `json:"type" validator:"required,ingredientType"`
}

func (cp CreateProductRequest) ToCommand() *command.CreateProductCommand {
	return command.NewCreateProductCommand(
		cp.Name, cp.Description, cp.Category, cp.Menu, toIngredients(cp.Ingredients),
	)
}

func toIngredients(ingredients []IngredientRequest) []command.Ingredient {
	var ingredientsCmd []command.Ingredient
	for _, ingredient := range ingredients {
		ingredientsCmd = append(ingredientsCmd, command.Ingredient{
			ID:       ingredient.ID,
			Name:     ingredient.Name,
			Quantity: ingredient.Quantity,
			Amount:   ingredient.Amount,
			Type:     command.GetIngredientTypeByName(ingredient.Type),
		})
	}
	return ingredientsCmd
}
