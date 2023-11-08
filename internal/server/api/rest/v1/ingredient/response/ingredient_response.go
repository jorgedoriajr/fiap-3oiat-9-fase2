package response

import "hamburgueria/internal/modules/ingredient/usecase/result"

type IngredientResponse struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

func FromFindIngredientsResult(result []result.FindIngredientResult) []IngredientResponse {
	var ingredients []IngredientResponse
	for _, ingredient := range result {
		ingredients = append(ingredients, IngredientResponse{
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
			Type:   ingredient.Type,
		})
	}
	return ingredients
}

func FromFindIngredientResult(ingredient result.FindIngredientResult) IngredientResponse {
	return IngredientResponse{
		Name:   ingredient.Name,
		Amount: ingredient.Amount,
		Type:   ingredient.Type,
	}
}

func FromCreateIngredientResult(result result.CreateIngredientResult) IngredientResponse {
	return IngredientResponse{
		Name:   result.Name,
		Amount: result.Amount,
		Type:   result.Type,
	}
}
