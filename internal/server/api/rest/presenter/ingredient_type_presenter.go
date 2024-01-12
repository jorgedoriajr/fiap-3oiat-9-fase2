package presenter

import (
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"hamburgueria/internal/server/api/rest/v1/ingredienttype/response"
)

func IngredientTypeResponseFromResult(result []result.IngredientTypeResult) []response.IngredientTypeResponse {
	var ingredientTypeResponse []response.IngredientTypeResponse
	for _, ingredientType := range result {
		ingredientTypeResponse = append(ingredientTypeResponse, response.IngredientTypeResponse{Name: ingredientType.Name})
	}
	return ingredientTypeResponse
}
