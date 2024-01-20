package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/payment/usecase/result"
	orderMocks "hamburgueria/tests/mocks/modules/order/port/output"
	mocks "hamburgueria/tests/mocks/modules/payment/port/input"
	"testing"
)

func TestProcessPaymentUseCase(t *testing.T) {

	t.Run(`should process payment`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		createPaymentUseCaseMock := mocks.NewCreatePaymentPort(t)
		processPaymentUseCase := ProcessPaymentUseCase{
			orderPersistenceGateway: orderPersistenceMock,
			createPaymentUseCase:    createPaymentUseCaseMock,
		}

		order := domain.Order{}

		payment := result.PaymentProcessed{
			PaymentId:   uuid.New(),
			PaymentData: "mocked",
		}

		createPaymentUseCaseMock.On("CreatePayment", mock.Anything, mock.Anything).Return(&payment, nil)

		orderPersistenceMock.On("Update", mock.Anything, mock.MatchedBy(func(c domain.Order) bool {
			return c.PaymentId == payment.PaymentId &&
				c.Status == valueobject.PaymentCreated &&
				len(c.History) == 1 &&
				c.History[0].Status == valueobject.PaymentCreated
		})).Return(nil)

		paymentCreated, err := processPaymentUseCase.ProcessPayment(context.TODO(), order)

		assert.Nil(t, err)
		assert.Equal(t, paymentCreated.PaymentData, "mocked")

		createPaymentUseCaseMock.AssertExpectations(t)
		createPaymentUseCaseMock.AssertCalled(t, "CreatePayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when create payment failed`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		createPaymentUseCaseMock := mocks.NewCreatePaymentPort(t)
		processPaymentUseCase := ProcessPaymentUseCase{
			orderPersistenceGateway: orderPersistenceMock,
			createPaymentUseCase:    createPaymentUseCaseMock,
		}

		order := domain.Order{}

		createPaymentUseCaseMock.On("CreatePayment", mock.Anything, mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		paymentCreated, err := processPaymentUseCase.ProcessPayment(context.TODO(), order)

		assert.NotNil(t, err)
		assert.Nil(t, paymentCreated)

		createPaymentUseCaseMock.AssertExpectations(t)
		createPaymentUseCaseMock.AssertCalled(t, "CreatePayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when update failed`, func(t *testing.T) {
		orderPersistenceMock := orderMocks.NewOrderPersistencePort(t)
		createPaymentUseCaseMock := mocks.NewCreatePaymentPort(t)
		processPaymentUseCase := ProcessPaymentUseCase{
			orderPersistenceGateway: orderPersistenceMock,
			createPaymentUseCase:    createPaymentUseCaseMock,
		}

		order := domain.Order{}

		payment := result.PaymentProcessed{
			PaymentId:   uuid.New(),
			PaymentData: "mocked",
		}

		createPaymentUseCaseMock.On("CreatePayment", mock.Anything, mock.Anything).Return(&payment, nil)

		orderPersistenceMock.On("Update", mock.Anything, mock.MatchedBy(func(c domain.Order) bool {
			return c.PaymentId == payment.PaymentId &&
				c.Status == valueobject.PaymentCreated &&
				len(c.History) == 1 &&
				c.History[0].Status == valueobject.PaymentCreated
		})).Return(errors.New("SOME_ERROR"))

		paymentCreated, err := processPaymentUseCase.ProcessPayment(context.TODO(), order)

		assert.NotNil(t, err)
		assert.Nil(t, paymentCreated)

		createPaymentUseCaseMock.AssertExpectations(t)
		createPaymentUseCaseMock.AssertCalled(t, "CreatePayment", mock.Anything, mock.Anything)

		orderPersistenceMock.AssertExpectations(t)
		orderPersistenceMock.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	})
}
