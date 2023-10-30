package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type CreateProductCategoryResult struct {
	ID   uuid.UUID
	Name string
}

func ToCreateProductCategoryResultFrom(entity entity.ProductCategoryEntity) CreateProductCategoryResult {
	return CreateProductCategoryResult{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
