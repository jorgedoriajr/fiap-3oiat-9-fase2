package write

import (
	"hamburgueria/internal/modules/product/domain/entity"
)

type InsertProductIngredientRWQueryMapper struct {
	Id           string `position:"0"`
	ProductId    string `position:"1"`
	IngredientId string `position:"2"`
	Quantity     int    `position:"3"`
	Amount       int    `position:"4"`
}

func ToInsertProductIngredientQueryMapper(product entity.ProductIngredientEntity) InsertProductIngredientRWQueryMapper {
	return InsertProductIngredientRWQueryMapper{
		Id:           product.ID.String(),
		ProductId:    product.ProductId.String(),
		IngredientId: product.IngredientId.String(),
		Quantity:     product.Quantity,
		Amount:       product.Amount,
	}
}
