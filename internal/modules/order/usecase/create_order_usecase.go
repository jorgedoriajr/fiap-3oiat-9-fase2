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
	"sync"
	"time"
)

type CreateOrderUseCase struct {
	productUseCase          productInputPort.CreateProductUseCasePort
	productPersistence      productPort.ProductPersistencePort
	orderPersistence        output.OrderPersistencePort
	orderHistoryPersistence output.OrderHistoryPersistencePort
	orderProductPersistence output.OrderProductPersistencePort
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
			product, err := c.productPersistence.GetByID(ctx, createProductCommand.Id)
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
			productCreated, err := c.createProduct(ctx, createProductCommand)
			if err != nil {
				return nil, err
			}
			productAmount = productCreated.Amount * createProductCommand.Quantity
			products = append(products, entity.OrderProduct{
				Id:        uuid.New(),
				ProductId: productCreated.Id,
				OrderId:   orderId,
				Quantity:  createProductCommand.Quantity,
				Amount:    productAmount,
			})
		}

		amount = amount + productAmount

	}

	order := entity.Order{
		Id:         orderId,
		CustomerId: createOrderCommand.CustomerDocument,
		Products:   products,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     "CREATED",
		Amount:     amount,
	}

	err := c.createOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &result.CreateOrderResult{
		Amount:      amount,
		PaymentData: "not implemented",
	}, err
}

// TODO need to create transaction
func (c CreateOrderUseCase) createOrder(ctx context.Context, order entity.Order) error {
	err := c.orderPersistence.Create(ctx, order)
	if err != nil {
		return err
	}

	err = c.createOrderHistory(ctx, order)
	if err != nil {
		return err
	}

	err = c.createOrderProducts(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (c CreateOrderUseCase) createOrderProducts(ctx context.Context, order entity.Order) error {
	for _, product := range order.Products {
		err := c.orderProductPersistence.Create(ctx, entity.OrderProduct{
			Id:        uuid.New(),
			ProductId: product.ProductId,
			OrderId:   order.Id,
			Quantity:  product.Quantity,
			Amount:    product.Amount,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c CreateOrderUseCase) createOrderHistory(ctx context.Context, order entity.Order) error {
	return c.orderHistoryPersistence.Create(ctx, entity.OrderHistory{
		Id:        uuid.New(),
		OrderId:   order.Id,
		Status:    order.Status,
		ChangeBy:  "user",
		CreatedAt: order.CreatedAt,
	})
}

func (c CreateOrderUseCase) createProduct(
	ctx context.Context,
	createOrderProductsCommand command.CreateOrderProductsCommand,
) (productResult.CreateProductResult, error) {

	var ingredients []productCommand.Ingredient

	for _, ingredient := range createOrderProductsCommand.Ingredients {
		ingredients = append(ingredients, productCommand.Ingredient{
			ID:       ingredient.Id.String(),
			Quantity: ingredient.Quantity,
		})
	}

	return c.productUseCase.AddProduct(ctx, productCommand.CreateProductCommand{
		Name:        "Personalized Product",
		Description: "Produto personalidado pelo cliente",
		Category:    createOrderProductsCommand.ProductCategory,
		Menu:        false,
		Ingredients: ingredients,
	})
}

var (
	createOrderUseCaseInstance CreateOrderUseCase
	createOrderUseCaseOnce     sync.Once
)

func GetCreateOrderUseCase(
	ProductUseCase productInputPort.CreateProductUseCasePort,
	ProductPersistence productPort.ProductPersistencePort,
	OrderPersistence output.OrderPersistencePort,
	OrderHistoryPersistence output.OrderHistoryPersistencePort,
	OrderProductPersistence output.OrderProductPersistencePort,
) CreateOrderUseCase {
	createOrderUseCaseOnce.Do(func() {
		createOrderUseCaseInstance = CreateOrderUseCase{
			productUseCase:          ProductUseCase,
			productPersistence:      ProductPersistence,
			orderPersistence:        OrderPersistence,
			orderHistoryPersistence: OrderHistoryPersistence,
			orderProductPersistence: OrderProductPersistence,
		}
	})
	return createOrderUseCaseInstance
}
