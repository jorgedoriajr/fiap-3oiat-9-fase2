package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type CreateProductCategoryCommand struct {
	Name string
}

func NewCreateProductCategoryCommand(
	Name string,
) *CreateProductCategoryCommand {

	cmd := &CreateProductCategoryCommand{
		Name: Name,
	}
	return cmd
}

func (pc CreateProductCategoryCommand) ToEntity() entity.ProductCategoryEntity {
	return entity.ProductCategoryEntity{
		ID:   uuid.New(),
		Name: pc.Name,
	}
}
