package entity

import "hamburgueria/internal/modules/product/domain/valueobject"

type ProductCategoryEntity struct {
	ID   int
	Name valueobject.ProductCategory
}
