package entity

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/valueobject"
	"time"
)

type ProductEntity struct {
	ID          uuid.UUID
	Number      int
	Name        string
	Amount      int
	Description string
	Category    valueobject.ProductCategory
	Menu        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
