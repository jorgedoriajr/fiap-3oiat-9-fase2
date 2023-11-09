package response

import "hamburgueria/internal/modules/product/usecase/result"

type FindProductWithIngredients struct {
	Name        string                    `json:"name"`
	Number      int                       `json:"number"`
	Amount      int                       `json:"amount"`
	Description string                    `json:"description"`
	Category    string                    `json:"category"`
	ImgPath     string                    `json:"imgPath"`
	Ingredients []FindProductsIngredients `json:"ingredients"`
}

type FindProductsIngredients struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity"`
}

func FromResultList(productResult []result.FindProductResult) []FindProductWithIngredients {
	var productsResponse []FindProductWithIngredients
	for _, product := range productResult {
		productsResponse = append(productsResponse, FromResult(product))
	}
	return productsResponse
}

func FromResult(product result.FindProductResult) FindProductWithIngredients {
	var ingredientsResponse []FindProductsIngredients
	for _, ingredient := range product.Ingredients {
		ingredientsResponse = append(ingredientsResponse, FindProductsIngredients{
			Number:   ingredient.Number,
			Name:     ingredient.Name,
			Amount:   ingredient.Amount,
			Quantity: ingredient.Quantity,
		})
	}
	return FindProductWithIngredients{
		Name:        product.Name,
		Number:      product.Number,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    product.Category,
		ImgPath:     product.ImgPath,
		Ingredients: ingredientsResponse,
	}
}
