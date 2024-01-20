package presenter

import (
	"hamburgueria/internal/modules/product/usecase/result"
	"hamburgueria/internal/web/api/rest/v1/product/response"
)

func ProductCreatedResponseFromResult(result result.ProductResult) response.ProductCreatedResponse {
	return response.ProductCreatedResponse{
		Id:          result.Id,
		Name:        result.Name,
		Amount:      result.Amount,
		Description: result.Description,
		Category:    result.Category,
		Menu:        result.Menu,
		ImgPath:     result.ImgPath,
	}
}

func ProductsResponseFromResultList(productResult []result.FindProductResult) []response.FindProductWithIngredients {
	var productsResponse []response.FindProductWithIngredients
	for _, product := range productResult {
		productsResponse = append(productsResponse, ProductResponseFromResult(product))
	}
	return productsResponse
}

func ProductResponseFromResult(product result.FindProductResult) response.FindProductWithIngredients {
	var ingredientsResponse []response.FindProductsIngredients
	for _, ingredient := range product.Ingredients {
		ingredientsResponse = append(ingredientsResponse, response.FindProductsIngredients{
			Number:   ingredient.Number,
			Name:     ingredient.Name,
			Amount:   ingredient.Amount,
			Quantity: ingredient.Quantity,
		})
	}
	return response.FindProductWithIngredients{
		Name:        product.Name,
		Number:      product.Number,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    product.Category,
		ImgPath:     product.ImgPath,
		Ingredients: ingredientsResponse,
	}
}
