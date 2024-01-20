package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	ingredientDomain "hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/order/domain"
	productDomain "hamburgueria/internal/modules/product/domain"
	mocks "hamburgueria/tests/mocks/modules/order/port/output"
	"testing"
	"time"
)

func TestListOrderUseCase(t *testing.T) {

	t.Run(`should find all orders`, func(t *testing.T) {
		orderPersistenceMock := mocks.NewOrderPersistencePort(t)
		listOrderUseCase := ListOrderUseCase{
			orderPersistenceGateway: orderPersistenceMock,
		}
		orderId := uuid.New()
		order := domain.Order{
			Id:         orderId,
			CustomerId: "Document",
			Products: []domain.OrderProduct{
				{
					Id: uuid.New(),
					Product: productDomain.Product{
						ID:          uuid.New(),
						Number:      1,
						Name:        "Product",
						Amount:      1000,
						Description: "Product Description",
						Category: productDomain.ProductCategory{
							Name: "Category",
						},
						Menu:      true,
						ImgPath:   "https://imgpath.com",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Ingredients: []productDomain.ProductIngredient{
							{
								ID:        uuid.UUID{},
								ProductId: uuid.UUID{},
								Ingredient: ingredientDomain.Ingredient{
									ID:     uuid.New(),
									Number: 1,
									Name:   "Ingredient",
									Amount: 1000,
									Type:   ingredientDomain.IngredientType{Name: "Type"},
								},
								Quantity: 1,
								Amount:   1000,
							},
						},
						Active: true,
					},
					OrderId:  orderId,
					Quantity: 1,
					Amount:   1000,
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Status:    "WAITING_PAYMENT",
			Amount:    1000,
			PaymentId: uuid.New(),
		}

		orderPersistenceMock.On("FindAll", mock.Anything).Return([]domain.Order{order}, nil)

		orders, err := listOrderUseCase.FindAllOrders(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(orders), 1)
		assert.Equal(t, orders[0].OrderId, orderId)
		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "FindAll", mock.Anything)
	})

	t.Run(`should return error when try to find all orders`, func(t *testing.T) {
		orderPersistenceMock := mocks.NewOrderPersistencePort(t)
		listOrderUseCase := ListOrderUseCase{
			orderPersistenceGateway: orderPersistenceMock,
		}

		orderPersistenceMock.On("FindAll", mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		orders, err := listOrderUseCase.FindAllOrders(context.TODO())

		assert.NotNil(t, err)
		assert.Nil(t, orders)
		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "FindAll", mock.Anything)
	})

	t.Run(`should find all orders`, func(t *testing.T) {
		orderPersistenceMock := mocks.NewOrderPersistencePort(t)
		listOrderUseCase := ListOrderUseCase{
			orderPersistenceGateway: orderPersistenceMock,
		}
		orderId := uuid.New()
		order := domain.Order{
			Id:         orderId,
			CustomerId: "Document",
			Products: []domain.OrderProduct{
				{
					Id: uuid.New(),
					Product: productDomain.Product{
						ID:          uuid.New(),
						Number:      1,
						Name:        "Product",
						Amount:      1000,
						Description: "Product Description",
						Category: productDomain.ProductCategory{
							Name: "Category",
						},
						Menu:      true,
						ImgPath:   "https://imgpath.com",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Ingredients: []productDomain.ProductIngredient{
							{
								ID:        uuid.UUID{},
								ProductId: uuid.UUID{},
								Ingredient: ingredientDomain.Ingredient{
									ID:     uuid.New(),
									Number: 1,
									Name:   "Ingredient",
									Amount: 1000,
									Type:   ingredientDomain.IngredientType{Name: "Type"},
								},
								Quantity: 1,
								Amount:   1000,
							},
						},
						Active: true,
					},
					OrderId:  orderId,
					Quantity: 1,
					Amount:   1000,
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Status:    "WAITING_PAYMENT",
			Amount:    1000,
			PaymentId: uuid.New(),
		}

		orderPersistenceMock.On("FindByStatus", mock.Anything, "WAITING_PAYMENT").Return([]domain.Order{order}, nil)

		orders, err := listOrderUseCase.FindByStatus(context.TODO(), "WAITING_PAYMENT")

		assert.Nil(t, err)
		assert.Equal(t, len(orders), 1)
		assert.Equal(t, orders[0].OrderId, orderId)
		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "FindByStatus", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when try to find all orders`, func(t *testing.T) {
		orderPersistenceMock := mocks.NewOrderPersistencePort(t)
		listOrderUseCase := ListOrderUseCase{
			orderPersistenceGateway: orderPersistenceMock,
		}

		orderPersistenceMock.On("FindByStatus", mock.Anything, "WAITING_PAYMENT").Return(nil, errors.New("SOME_ERROR"))

		orders, err := listOrderUseCase.FindByStatus(context.TODO(), "WAITING_PAYMENT")

		assert.NotNil(t, err)
		assert.Nil(t, orders)
		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "FindByStatus", mock.Anything, mock.Anything)
	})
}
