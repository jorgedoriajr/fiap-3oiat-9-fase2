package domain

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Number      int
	Name        string
	Amount      int
	Description string
	Category    ProductCategory
	Menu        bool
	ImgPath     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Ingredients []ProductIngredient
	Active      bool
}

type ProductIngredient struct {
	ID         uuid.UUID
	ProductId  uuid.UUID
	Ingredient domain.Ingredient
	Quantity   int
	Amount     int
}
