package result

import (
	uuid "github.com/vgarvardt/pgx-google-uuid/v5"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/domain/valueobject"
	"time"
)

type CreateProductResult struct {
	ID          *uuid.UUID
	Number      *int
	Name        string
	Amount      int
	Description string
	Category    valueobject.ProductCategory
	Menu        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToCreateProductResultFrom(entity entity.ProductEntity) CreateProductResult {
	return CreateProductResult{
		Name:        entity.Name,
		Amount:      entity.Amount,
		Description: entity.Description,
		Category:    entity.Category,
		Menu:        entity.Menu,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
