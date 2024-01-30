package result

import (
	"hamburgueria/internal/modules/product/domain"
)

type FindProductResult struct {
	Name        string
	Number      int
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
	Ingredients []ProductIngredientsResult
	Active      bool
}

type ProductIngredientsResult struct {
	Number   int
	Name     string
	Amount   int
	Type     string
	Quantity int
}

func FromProductDomain(product domain.Product) FindProductResult {
	var ingredients []ProductIngredientsResult

	for _, productIngredient := range product.Ingredients {
		ingredients = append(ingredients, FromProductIngredientDomain(productIngredient))
	}

	return FindProductResult{
		Name:        product.Name,
		Number:      product.Number,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    product.Category.Name,
		Menu:        product.Menu,
		ImgPath:     product.ImgPath,
		Ingredients: ingredients,
		Active:      product.Active,
	}
}

func FromProductIngredientDomain(productIngredient domain.ProductIngredient) ProductIngredientsResult {
	return ProductIngredientsResult{
		Number:   productIngredient.Ingredient.Number,
		Name:     productIngredient.Ingredient.Name,
		Amount:   productIngredient.Amount,
		Type:     productIngredient.Ingredient.Type.Name,
		Quantity: productIngredient.Quantity,
	}
}
