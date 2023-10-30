package write

import (
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type InsertProductRWQueryMapper struct {
	Name        string    `position:"0"`
	Amount      int       `position:"1"`
	Description string    `position:"2"`
	Category    string    `position:"3"`
	Menu        bool      `position:"4"`
	CreatedAt   time.Time `position:"5"`
	UpdatedAt   time.Time `position:"6"`
}

func ToInsertProductQueryMapper(product entity.ProductEntity) InsertProductRWQueryMapper {
	return InsertProductRWQueryMapper{
		Name:        product.Name,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    string(product.Category),
		Menu:        product.Menu,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
