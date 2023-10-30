package write

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type InsertProductCategoryRWQueryMapper struct {
	Id   uuid.UUID `position:"0"`
	Name string    `position:"1"`
}

func ToInsertProductCategoryQueryMapper(productCategory entity.ProductCategoryEntity) InsertProductCategoryRWQueryMapper {
	return InsertProductCategoryRWQueryMapper{
		Id:   productCategory.ID,
		Name: productCategory.Name,
	}
}
