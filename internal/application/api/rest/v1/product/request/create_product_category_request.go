package request

import (
	"hamburgueria/internal/modules/product/usecase/command"
)

type CreateProductCategoryRequest struct {
	Name string `json:"name" validator:"required"`
}

func (cp CreateProductCategoryRequest) ToCommand() command.CreateProductCategoryCommand {
	return *command.NewCreateProductCategoryCommand(cp.Name)
}
