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
	ImgPath     string    `position:"6"`
	CreatedAt   time.Time `position:"7"`
	UpdatedAt   time.Time `position:"8"`
}

func ToInsertProductQueryMapper(product entity.ProductEntity) InsertProductRWQueryMapper {
	return InsertProductRWQueryMapper{
		Id:          product.ID.String(),
		Name:        product.Name,
		Amount:      product.Amount,
		Description: product.Description,
		Category:    product.Category,
		Menu:        product.Menu,
		ImgPath:     product.ImgPath,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

type UpdateProductRWQueryMapper struct {
	Number      int       `position:"0"`
	Name        *string   `position:"1"`
	Description *string   `position:"2"`
	Category    *string   `position:"3"`
	Menu        *bool     `position:"4"`
	ImgPath     *string   `position:"5"`
	UpdatedAt   time.Time `position:"6"`
}

func ToUpdateProductQueryMapper(product entity.ProductEntity) UpdateProductRWQueryMapper {
	return UpdateProductRWQueryMapper{
		Number:      product.Number,
		Name:        &product.Name,
		Description: &product.Description,
		Category:    &product.Category,
		Menu:        &product.Menu,
		ImgPath:     &product.ImgPath,
		UpdatedAt:   product.UpdatedAt,
	}
}