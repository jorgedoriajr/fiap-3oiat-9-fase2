package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type FindProductCategoryQueryResult struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (fc FindProductQueryResult) ToProductCategoryEntity() *entity.ProductCategoryEntity {
	return &entity.ProductCategoryEntity{
		ID:   fc.ID,
		Name: fc.Name,
	}
}

func (fc FindProductQueryResult) ToProductCategoryCommandResult() *entity.ProductCategoryEntity {
	return &entity.ProductCategoryEntity{
		ID:   fc.ID,
		Name: fc.Name,
	}
}
