package usecase

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/domain/response"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/command"
	productEntity "hamburgueria/internal/modules/product/domain/entity"
	productPort "hamburgueria/internal/modules/product/ports/output"
	productCommand "hamburgueria/internal/modules/product/usecase/command"
	"time"
)

type CreateOrderUseCase struct {
	ProductUseCase     productPort.CreateProductUseCasePort
	ProductPersistence productPort.ProductPersistencePort
	OrderPersistence   output.OrderPersistencePort
}

func (c CreateOrderUseCase) AddOrder(
	ctx context.Context,
	createOrderCommand command.CreateOrderCommand,
) (*response.OrderResponse, error) {

	var amount int64

	for _, createProductCommand := range createOrderCommand.Products {

		product, err := c.findProductByType(ctx, createProductCommand)
		if err != nil {
			return nil, err
		}

		if product == nil {
			_, err := c.ProductUseCase.AddProduct(ctx, productCommand.CreateProductCommand{
				Name:        "Personalized Product",
				Description: "Produto personalidado pelo cliente",
				Category:    "Dish",
				Menu:        false,
				Ingredients: nil,
			})
			if err != nil {
				return nil, err
			}
		}

		amount = amount + int64(product.Amount*createProductCommand.Quantity)

	}

	err := c.OrderPersistence.Create(
		ctx,
		entity.Order{
			Id:         uuid.UUID{},
			CustomerId: createOrderCommand.CustomerDocument,
			Products:   nil,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Status:     "CREATED",
			Amount:     amount,
		},
	)
	if err != nil {
		return nil, err
	}

	return &response.OrderResponse{
		Amount:      amount,
		PaymentData: "",
	}, err
}

func (c CreateOrderUseCase) findProductByType(ctx context.Context, createProductCommand command.CreateOrderProductsCommand) (*productEntity.ProductEntity, error) {
	var product productEntity.ProductEntity
	var err error

	if createProductCommand.Type == "default" {
		product, err = c.ProductPersistence.GetByID(ctx, createProductCommand.Id.String())
		/*if product == nil { //TODO vai voltar nulo?
			return errors.New("product not found")
		}*/
	}

	if createProductCommand.Type == "personalized" {
		product, err = c.ProductPersistence.GetByID(ctx, createProductCommand.Id.String()) // TODO findByIngredients
	}
	return &product, err
}
