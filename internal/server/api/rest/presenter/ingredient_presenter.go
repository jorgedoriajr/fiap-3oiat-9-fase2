package presenter

import (
	result2 "hamburgueria/internal/modules/ingredient/usecase/result"
	"hamburgueria/internal/server/api/rest/v1/ingredient/response"
)

func FindIngredientsResponseFromResult(result []result2.FindIngredientResult) []response.IngredientResponse {
	var ingredients []response.IngredientResponse
	for _, ingredient := range result {
		ingredients = append(ingredients, response.IngredientResponse{
			Number: ingredient.Number,
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
			Type:   ingredient.Type,
		})
	}
	return ingredients
}

func FindIngredientResponseFromResult(ingredient result2.FindIngredientResult) response.IngredientResponse {
	return response.IngredientResponse{
		Number: ingredient.Number,
		Name:   ingredient.Name,
		Amount: ingredient.Amount,
		Type:   ingredient.Type,
	}
}

func CreateIngredientResponseFromResult(result result2.CreateIngredientResult) response.IngredientResponse {
	return response.IngredientResponse{
		Name:   result.Name,
		Amount: result.Amount,
		Type:   result.Type,
	}
}
