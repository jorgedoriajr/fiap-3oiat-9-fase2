package entity

import (
	"hamburgueria/internal/modules/product/domain/valueobject"
	"time"
)

type ProductEntity struct {
	ID          int
	Name        string
	Amount      int
	Description string
	Category    valueobject.ProductCategory
	Menu        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
