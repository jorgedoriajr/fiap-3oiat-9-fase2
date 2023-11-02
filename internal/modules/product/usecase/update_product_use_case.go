package usecase

import (
	"context"
	"fmt"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
	"hamburgueria/pkg/querymapper"
)

type UpdateProductUseCase struct {
	productPersistencePort  output.ProductPersistencePort
	categoryPersistencePort output.ProductCategoryPersistencePort
}

func (u UpdateProductUseCase) UpdateProduct(ctx context.Context, command command.UpdateProductCommand) (result.UpdateProductResult, error) {

	updateQueryCommand := querymapper.UpdateQueryCommand{
		Table:      "product",
		Condition:  fmt.Sprintf("number = %d", command.Number),
		UpdateData: command,
	}

	err := u.productPersistencePort.Update(ctx, updateQueryCommand)
	if err != nil {
		return result.UpdateProductResult{}, err
	}

	productByNumber, err := u.productPersistencePort.GetByNumber(ctx, command.Number)
	if err != nil {
		return result.UpdateProductResult{}, err
	}

	return result.UpdateProductResultFromEntity(*productByNumber), nil
}

func NewUpdateProductUseCase(
	productPersistencePort output.ProductPersistencePort,
) *UpdateProductUseCase {
	//sync.OnceValue()
	return &UpdateProductUseCase{productPersistencePort: productPersistencePort}
}
