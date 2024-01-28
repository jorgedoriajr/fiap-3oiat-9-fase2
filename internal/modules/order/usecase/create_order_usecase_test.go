package usecase

import (
	"context"
	"errors"
	customerDomain "hamburgueria/internal/modules/customer/domain"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/usecase/command"
	"hamburgueria/internal/modules/order/usecase/result"
	productDomain "hamburgueria/internal/modules/product/domain"
	customerMocks "hamburgueria/tests/mocks/modules/customer/port/output"
	mocks "hamburgueria/tests/mocks/modules/order/port/input"
	orderMocks "hamburgueria/tests/mocks/modules/order/port/output"
	productMocks "hamburgueria/tests/mocks/modules/product/ports/output"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderUseCase(t *testing.T) {

	t.Run(`should create order`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		customerPersistenceMock := customerMocks.NewCustomerPersistencePort(t)
		productPersistenceMock := productMocks.NewProductPersistencePort(t)
		processPaymentUseCaseMock := mocks.NewProcessPaymentPort(t)
		createOrderUseCase := CreateOrderUseCase{
			customerPersistenceGateway: customerPersistenceMock,
			productPersistenceGateway:  productPersistenceMock,
			orderPersistenceGateway:    orderPersistenceMock,
			processPaymentUseCase:      processPaymentUseCaseMock,
		}

		createOrderCommand := command.CreateOrderCommand{
			CustomerDocument: "Document",
			Products: []command.CreateOrderProductsCommand{
				{
					Number:   1,
					Quantity: 2,
					Type:     "default",
				},
			},
		}

		product := productDomain.Product{
			ID:          uuid.New(),
			Number:      1,
			Name:        "Product",
			Amount:      1500,
			Description: "Product Description",
			Category: productDomain.ProductCategory{
				Name: "Category",
			},
			Menu:      true,
			ImgPath:   "https://imgPath.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Active:    true,
		}

		customerPersistenceMock.On("Get", mock.Anything, "Document").Return(&customerDomain.Customer{
			Document:       "Document",
			Name:           "Customer",
			Phone:          "Phone",
			Email:          "Mail",
			OptInPromotion: false,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}, nil)

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&product, nil)

		orderPersistenceMock.On("Create", mock.Anything, mock.MatchedBy(func(c domain.Order) bool {
			return c.Status == valueobject.Created &&
				c.Amount == 3000
		})).Return(nil)

		orderPersistenceMock.On("FindById", mock.Anything, mock.Anything).Return(
			&domain.Order{Number: 1},
			nil,
		)

		processPaymentUseCaseMock.On("ProcessPayment", mock.Anything, mock.Anything).Return(&result.PaymentCreatedResult{
			PaymentData: "mocked",
		}, nil)

		orderCreated, err := createOrderUseCase.AddOrder(context.TODO(), createOrderCommand)

		assert.Nil(t, err)
		assert.Equal(t, 3000, orderCreated.Amount)
		assert.Equal(t, orderCreated.Number, 1)
		assert.Equal(t, "mocked", orderCreated.PaymentData)

		customerPersistenceMock.AssertExpectations(t)
		customerPersistenceMock.AssertCalled(t, "Get", mock.Anything, "Document")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, 1)

		processPaymentUseCaseMock.AssertExpectations(t)
		processPaymentUseCaseMock.AssertCalled(t, "ProcessPayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when customer not found`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		customerPersistenceMock := customerMocks.NewCustomerPersistencePort(t)
		productPersistenceMock := productMocks.NewProductPersistencePort(t)
		processPaymentUseCaseMock := mocks.NewProcessPaymentPort(t)
		createOrderUseCase := CreateOrderUseCase{
			customerPersistenceGateway: customerPersistenceMock,
			productPersistenceGateway:  productPersistenceMock,
			orderPersistenceGateway:    orderPersistenceMock,
			processPaymentUseCase:      processPaymentUseCaseMock,
		}

		createOrderCommand := command.CreateOrderCommand{
			CustomerDocument: "Document",
			Products: []command.CreateOrderProductsCommand{
				{
					Number:   1,
					Quantity: 2,
					Type:     "default",
				},
			},
		}

		customerPersistenceMock.On("Get", mock.Anything, "Document").Return(nil, nil)

		orderCreated, err := createOrderUseCase.AddOrder(context.TODO(), createOrderCommand)

		assert.NotNil(t, err)
		assert.Nil(t, orderCreated)

		customerPersistenceMock.AssertExpectations(t)
		customerPersistenceMock.AssertCalled(t, "Get", mock.Anything, "Document")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertNotCalled(t, "GetByNumber", mock.Anything, 1)

		processPaymentUseCaseMock.AssertExpectations(t)
		processPaymentUseCaseMock.AssertNotCalled(t, "ProcessPayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when product not found`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		customerPersistenceMock := customerMocks.NewCustomerPersistencePort(t)
		productPersistenceMock := productMocks.NewProductPersistencePort(t)
		processPaymentUseCaseMock := mocks.NewProcessPaymentPort(t)
		createOrderUseCase := CreateOrderUseCase{
			customerPersistenceGateway: customerPersistenceMock,
			productPersistenceGateway:  productPersistenceMock,
			orderPersistenceGateway:    orderPersistenceMock,
			processPaymentUseCase:      processPaymentUseCaseMock,
		}

		createOrderCommand := command.CreateOrderCommand{
			CustomerDocument: "Document",
			Products: []command.CreateOrderProductsCommand{
				{
					Number:   1,
					Quantity: 2,
					Type:     "default",
				},
			},
		}

		customerPersistenceMock.On("Get", mock.Anything, "Document").Return(&customerDomain.Customer{
			Document:       "Document",
			Name:           "Customer",
			Phone:          "Phone",
			Email:          "Mail",
			OptInPromotion: false,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}, nil)

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(nil, nil)

		orderCreated, err := createOrderUseCase.AddOrder(context.TODO(), createOrderCommand)

		assert.NotNil(t, err)
		assert.Nil(t, orderCreated)

		customerPersistenceMock.AssertExpectations(t)
		customerPersistenceMock.AssertCalled(t, "Get", mock.Anything, "Document")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, 1)

		processPaymentUseCaseMock.AssertExpectations(t)
		processPaymentUseCaseMock.AssertNotCalled(t, "ProcessPayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when create failed`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		customerPersistenceMock := customerMocks.NewCustomerPersistencePort(t)
		productPersistenceMock := productMocks.NewProductPersistencePort(t)
		processPaymentUseCaseMock := mocks.NewProcessPaymentPort(t)
		createOrderUseCase := CreateOrderUseCase{
			customerPersistenceGateway: customerPersistenceMock,
			productPersistenceGateway:  productPersistenceMock,
			orderPersistenceGateway:    orderPersistenceMock,
			processPaymentUseCase:      processPaymentUseCaseMock,
		}

		createOrderCommand := command.CreateOrderCommand{
			CustomerDocument: "Document",
			Products: []command.CreateOrderProductsCommand{
				{
					Number:   1,
					Quantity: 2,
					Type:     "default",
				},
			},
		}

		customerPersistenceMock.On("Get", mock.Anything, "Document").Return(&customerDomain.Customer{
			Document:       "Document",
			Name:           "Customer",
			Phone:          "Phone",
			Email:          "Mail",
			OptInPromotion: false,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}, nil)

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&productDomain.Product{Amount: 1500}, nil)

		orderPersistenceMock.On("Create", mock.Anything, mock.MatchedBy(func(c domain.Order) bool {
			return c.Status == valueobject.Created &&
				c.Amount == 3000
		})).Return(errors.New("SOME_ERROR"))

		orderCreated, err := createOrderUseCase.AddOrder(context.TODO(), createOrderCommand)

		assert.NotNil(t, err)
		assert.Nil(t, orderCreated)

		customerPersistenceMock.AssertExpectations(t)
		customerPersistenceMock.AssertCalled(t, "Get", mock.Anything, "Document")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, 1)

		processPaymentUseCaseMock.AssertExpectations(t)
		processPaymentUseCaseMock.AssertNotCalled(t, "ProcessPayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when payment failed`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		customerPersistenceMock := customerMocks.NewCustomerPersistencePort(t)
		productPersistenceMock := productMocks.NewProductPersistencePort(t)
		processPaymentUseCaseMock := mocks.NewProcessPaymentPort(t)
		createOrderUseCase := CreateOrderUseCase{
			customerPersistenceGateway: customerPersistenceMock,
			productPersistenceGateway:  productPersistenceMock,
			orderPersistenceGateway:    orderPersistenceMock,
			processPaymentUseCase:      processPaymentUseCaseMock,
		}

		createOrderCommand := command.CreateOrderCommand{
			CustomerDocument: "Document",
			Products: []command.CreateOrderProductsCommand{
				{
					Number:   1,
					Quantity: 2,
					Type:     "default",
				},
			},
		}

		customerPersistenceMock.On("Get", mock.Anything, "Document").Return(&customerDomain.Customer{
			Document:       "Document",
			Name:           "Customer",
			Phone:          "Phone",
			Email:          "Mail",
			OptInPromotion: false,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}, nil)

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&productDomain.Product{Amount: 1500}, nil)

		orderPersistenceMock.On("Create", mock.Anything, mock.MatchedBy(func(c domain.Order) bool {
			return c.Status == valueobject.Created &&
				c.Amount == 3000
		})).Return(nil)

		processPaymentUseCaseMock.On("ProcessPayment", mock.Anything, mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		orderCreated, err := createOrderUseCase.AddOrder(context.TODO(), createOrderCommand)

		assert.NotNil(t, err)
		assert.Nil(t, orderCreated)

		customerPersistenceMock.AssertExpectations(t)
		customerPersistenceMock.AssertCalled(t, "Get", mock.Anything, "Document")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, 1)

		processPaymentUseCaseMock.AssertExpectations(t)
		processPaymentUseCaseMock.AssertCalled(t, "ProcessPayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})
}
