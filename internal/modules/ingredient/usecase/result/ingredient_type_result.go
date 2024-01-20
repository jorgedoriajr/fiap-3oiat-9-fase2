package result

import "hamburgueria/internal/web/api/rest/v1/ingredienttype/response"

type IngredientTypeResult struct {
	Name string
}

func (i IngredientTypeResult) ToResponse() response.IngredientTypeResponse {
	return response.IngredientTypeResponse{Name: i.Name}
}
