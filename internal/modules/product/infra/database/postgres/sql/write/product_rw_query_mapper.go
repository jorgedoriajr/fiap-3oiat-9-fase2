package write

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type InsertProductRWQueryMapper struct {
	ID          uuid.UUID   `position:"0"`
	Name        string      `position:"1"`
	Amount      int         `position:"2"`
	Description string      `position:"3"`
	Category    string      `position:"4"`
	Menu        bool        `position:"5"`
	Ingredients []uuid.UUID `position:"6"`
	CreatedAt   time.Time   `position:"7"`
	UpdatedAt   time.Time   `position:"8"`
}

func ToInsertProductQueryMapper(product entity.ProductEntity) InsertProductRWQueryMapper {
	return InsertProductRWQueryMapper{
		ID:          product.ID,
		Name:        product.Name,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    product.Category,
		Menu:        product.Menu,
		Ingredients: product.Ingredients,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
