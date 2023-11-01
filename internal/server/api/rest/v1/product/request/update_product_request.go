package request

import (
	"hamburgueria/internal/modules/product/usecase/command"
)

type UpdateProductRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Menu        *bool   `json:"menu"`
	ImgPath     *string `json:"imgPath"`
}

func (cp UpdateProductRequest) ToCommandWithNumber(number int) command.UpdateProductCommand {
	return command.NewUpdateProductCommand(
		number, cp.Name, cp.Description, nil, cp.Menu, cp.ImgPath,
	)
}
