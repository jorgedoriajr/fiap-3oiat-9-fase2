package model

import (
	"github.com/google/uuid"
	ingredientDomain "hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/ingredient/infra/database/model"
	"hamburgueria/internal/modules/product/domain"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Number      int `gorm:"autoIncrement:true;unique"`
	Name        string
	Amount      int
	Description string
	Category    ProductCategory `gorm:"foreignKey:Name"`
	Menu        bool
	ImgPath     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Ingredients []ProductIngredient
}

type ProductIngredient struct {
	ID         uuid.UUID
	Number     int
	ProductId  uuid.UUID
	Ingredient model.Ingredient
	Quantity   int
	Amount     int
}

func (p Product) ToDomain() *domain.Product {
	var ingredients []domain.ProductIngredient

	for _, ingredient := range p.Ingredients {
		ingredients = append(ingredients, ingredient.ToDomain())
	}

	return &domain.Product{
		ID:          p.ID,
		Number:      p.Number,
		Name:        p.Name,
		Amount:      p.Amount,
		Description: p.Description,
		Category: domain.ProductCategory{
			Name:         p.Category.Name,
			AcceptCustom: p.Category.AcceptCustom,
		},
		Menu:        p.Menu,
		ImgPath:     p.ImgPath,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Ingredients: ingredients,
	}
}

func ProductFromDomain(product domain.Product) Product {
	var ingredients []ProductIngredient
	for _, ingredient := range product.Ingredients {
		ingredients = append(ingredients, ProductIngredientFromDomain(ingredient))
	}
	return Product{
		ID:          product.ID,
		Number:      product.Number,
		Name:        product.Name,
		Amount:      product.Amount,
		Description: product.Description,
		Category: ProductCategory{
			Name:         product.Category.Name,
			AcceptCustom: product.Category.AcceptCustom,
		},
		Menu:        product.Menu,
		ImgPath:     product.ImgPath,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Ingredients: ingredients,
	}
}

func (pi ProductIngredient) ToDomain() domain.ProductIngredient {
	return domain.ProductIngredient{
		ID:        pi.ID,
		Number:    pi.Number,
		ProductId: pi.ProductId,
		Ingredient: ingredientDomain.Ingredient{
			ID:     pi.Ingredient.ID,
			Number: pi.Ingredient.Number,
			Name:   pi.Ingredient.Name,
			Amount: pi.Ingredient.Amount,
			Type:   pi.Ingredient.Type.Name,
		},
		Quantity: pi.Quantity,
		Amount:   pi.Amount,
	}
}

func ProductIngredientFromDomain(productIngredient domain.ProductIngredient) ProductIngredient {
	return ProductIngredient{
		ID:        productIngredient.ID,
		Number:    productIngredient.Number,
		ProductId: productIngredient.ProductId,
		Ingredient: model.Ingredient{
			ID:     productIngredient.Ingredient.ID,
			Number: productIngredient.Ingredient.Number,
			Name:   productIngredient.Ingredient.Name,
			Amount: productIngredient.Ingredient.Amount,
			Type:   model.IngredientType{Name: productIngredient.Ingredient.Type},
		},
		Quantity: productIngredient.Quantity,
		Amount:   productIngredient.Amount,
	}
}
