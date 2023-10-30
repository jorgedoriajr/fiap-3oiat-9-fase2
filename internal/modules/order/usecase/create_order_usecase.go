package usecase

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/command"
	"hamburgueria/internal/modules/order/usecase/result"
	productInputPort "hamburgueria/internal/modules/product/ports/input"
	productPort "hamburgueria/internal/modules/product/ports/output"
	productCommand "hamburgueria/internal/modules/product/usecase/command"
	productResult "hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type CreateOrderUseCase struct {
	ProductUseCase     productInputPort.CreateProductUseCasePort
	ProductPersistence productPort.ProductPersistencePort
	OrderPersistence   output.OrderPersistencePort
}

func (c CreateOrderUseCase) AddOrder(
	ctx context.Context,
	createOrderCommand command.CreateOrderCommand,
) (*result.CreateOrderResult, error) {

	var amount int
	var products []entity.OrderProduct
	orderId := uuid.New()

	for _, createProductCommand := range createOrderCommand.Products {
		var productAmount int
		if createProductCommand.Type == "default" {
			product, err := c.ProductPersistence.GetByID(ctx, createProductCommand.Id)
			if err != nil {
				return nil, err
			}
			productAmount = product.Amount * createProductCommand.Quantity
			products = append(products, entity.OrderProduct{
				Id:        uuid.New(),
				ProductId: product.ID,
				OrderId:   orderId,
				Quantity:  createProductCommand.Quantity,
				Amount:    productAmount,
			})
		} else {
			productCreated, err := c.createProduct(ctx)
			if err != nil {
				return nil, err
			}
			productAmount = productCreated.Amount * createProductCommand.Quantity
			products = append(products, entity.OrderProduct{
				Id:        uuid.New(),
				ProductId: productCreated.ID,
				OrderId:   orderId,
				Quantity:  createProductCommand.Quantity,
				Amount:    productAmount,
			})
		}

		amount = amount + productAmount

	}

	err := c.OrderPersistence.Create(
		ctx,
		entity.Order{
			Id:         uuid.New(),
			CustomerId: createOrderCommand.CustomerDocument,
			Products:   products,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Status:     "CREATED",
			Amount:     amount,
		},
	)
	if err != nil {
		return nil, err
	}

	return &result.CreateOrderResult{
		Amount:      amount,
		PaymentData: "not implemented",
	}, err
}

func (c CreateOrderUseCase) createProduct(ctx context.Context) (productResult.CreateProductResult, error) {
	return c.ProductUseCase.AddProduct(ctx, productCommand.CreateProductCommand{
		Name:        "Personalized Product",
		Description: "Produto personalidado pelo cliente",
		Category:    "Dish",
		Menu:        false,
		Ingredients: nil, //TODO fix
	})
}
