package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type CreateProductResult struct {
	Id          uuid.UUID
	Number      int
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
		Id:          entity.ID,
		Number:      entity.Number,
		Name:        entity.Name,
		Amount:      entity.Amount,
		Description: entity.Description,
		Category:    entity.Category,
		Menu:        entity.Menu,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
