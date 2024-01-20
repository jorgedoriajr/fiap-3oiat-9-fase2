package request

import (
	"hamburgueria/internal/modules/product/usecase/command"
)

type UpdateProductRequest struct {
	Number      int                       `json:"number"`
	Name        *string                   `json:"name"`
	Description *string                   `json:"description"`
	Category    *string                   `json:"category"`
	Menu        *bool                     `json:"menu"`
	ImgPath     *string                   `json:"imgPath"`
	Ingredients []UpdateIngredientRequest `json:"ingredients"`
}

type UpdateIngredientRequest struct {
	Number   int `json:"number" validator:"required"`
	Quantity int `json:"quantity" validator:"required"`
}

func (cp UpdateProductRequest) ToCommand() command.UpdateProductCommand {
	return *command.NewUpdateProductCommand(
		cp.Number,
		cp.Name,
		cp.Description,
		cp.Category,
		cp.Menu,
		toUpdateIngredientsRequest(cp.Ingredients),
		cp.ImgPath,
	)
}

func toUpdateIngredientsRequest(ingredients []UpdateIngredientRequest) []command.Ingredient {
	var ingredientsCmd []command.Ingredient
	for _, ingredient := range ingredients {
		ingredientsCmd = append(ingredientsCmd, command.Ingredient{
			Number:   ingredient.Number,
			Quantity: ingredient.Quantity,
		})
	}
	return ingredientsCmd
}
