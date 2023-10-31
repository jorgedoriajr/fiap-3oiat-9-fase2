package write

import (
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type InsertProductRWQueryMapper struct {
	Id          string    `position:"0"`
	Name        string    `position:"1"`
	Amount      int       `position:"2"`
	Description string    `position:"3"`
	Category    string    `position:"4"`
	Menu        bool      `position:"5"`
	CreatedAt   time.Time `position:"6"`
	UpdatedAt   time.Time `position:"7"`
}

func ToInsertProductQueryMapper(product entity.ProductEntity) InsertProductRWQueryMapper {
	return InsertProductRWQueryMapper{
		Id:          product.ID.String(),
		Name:        product.Name,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    product.Category,
		Menu:        product.Menu,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
