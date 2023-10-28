package result

import (
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type CreateProductResult struct {
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToCreateProductResultFrom(entity entity.ProductEntity) CreateProductResult {
	return CreateProductResult{
		Name:        entity.Name,
		Amount:      entity.Amount,
		Description: entity.Description,
		Category:    string(entity.Category),
		Menu:        entity.Menu,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
