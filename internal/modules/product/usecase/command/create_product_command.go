package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type CreateProductCommand struct {
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	Ingredients []uuid.UUID
}

func (cmd CreateProductCommand) ToProductEntity() entity.ProductEntity {
	return entity.ProductEntity{}
}
