package read

import (
	"hamburgueria/internal/modules/product/domain/entity"
)

type FindProductCategoryQueryResult struct {
	AcceptCustom bool   `db:"accept_custom"`
	Name         string `db:"name"`
}

func (fc FindProductCategoryQueryResult) ToProductCategoryEntity() *entity.ProductCategoryEntity {
	return &entity.ProductCategoryEntity{
		Name:         fc.Name,
		AcceptCustom: fc.AcceptCustom,
	}
}

func (fc FindProductCategoryQueryResult) ToEntity() entity.ProductCategoryEntity {
	return entity.ProductCategoryEntity{
		Name:         fc.Name,
		AcceptCustom: fc.AcceptCustom,
	}
}

func ToProductCategoryEntityList(results []FindProductCategoryQueryResult) []entity.ProductCategoryEntity {
	entities := make([]entity.ProductCategoryEntity, len(results))

	for i, result := range results {
		entities[i] = result.ToEntity()
	}

	return entities
}
