package usecase

import (
	"context"
	"errors"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/payment/usecase/result"
	orderMocks "hamburgueria/tests/mocks/modules/order/port/input"
	mocks "hamburgueria/tests/mocks/modules/payment/port/input"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProcessPaymentUseCase(t *testing.T) {

	t.Run(`should process payment`, func(t *testing.T) {
		updateOrderUseCaseMock := orderMocks.NewUpdateOrderPort(t)
		createPaymentUseCaseMock := mocks.NewCreatePaymentPort(t)
		processPaymentUseCase := ProcessPaymentUseCase{
			updateOrderUseCase:   updateOrderUseCaseMock,
			createPaymentUseCase: createPaymentUseCaseMock,
		}

		order := domain.Order{}

		payment := result.PaymentProcessed{
			PaymentId:   uuid.New(),
			PaymentData: "mocked",
		}

		createPaymentUseCaseMock.On("CreatePayment", mock.Anything, mock.Anything).Return(&payment, nil)

		updateOrderUseCaseMock.On("Update", mock.Anything, payment.OrderId, valueobject.PaymentCreated, &payment.PaymentId).Return(nil)

		paymentCreated, err := processPaymentUseCase.ProcessPayment(context.TODO(), order)

		assert.Nil(t, err)
		assert.Equal(t, paymentCreated.PaymentData, "mocked")

		createPaymentUseCaseMock.AssertExpectations(t)
		createPaymentUseCaseMock.AssertCalled(t, "CreatePayment", mock.Anything, mock.Anything)

		updateOrderUseCaseMock.AssertExpectations(t)
		updateOrderUseCaseMock.AssertCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run(`should return error when create payment failed`, func(t *testing.T) {
		updateOrderUseCaseMock := orderMocks.NewUpdateOrderPort(t)
		createPaymentUseCaseMock := mocks.NewCreatePaymentPort(t)
		processPaymentUseCase := ProcessPaymentUseCase{
			updateOrderUseCase:   updateOrderUseCaseMock,
			createPaymentUseCase: createPaymentUseCaseMock,
		}

		order := domain.Order{}

		createPaymentUseCaseMock.On("CreatePayment", mock.Anything, mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		paymentCreated, err := processPaymentUseCase.ProcessPayment(context.TODO(), order)

		assert.NotNil(t, err)
		assert.Nil(t, paymentCreated)

		createPaymentUseCaseMock.AssertExpectations(t)
		createPaymentUseCaseMock.AssertCalled(t, "CreatePayment", mock.Anything, mock.Anything)

		updateOrderUseCaseMock.AssertExpectations(t)
		updateOrderUseCaseMock.AssertNotCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run(`should return error when update failed`, func(t *testing.T) {
		updateOrderUseCaseMock := orderMocks.NewUpdateOrderPort(t)
		createPaymentUseCaseMock := mocks.NewCreatePaymentPort(t)
		processPaymentUseCase := ProcessPaymentUseCase{
			updateOrderUseCase:   updateOrderUseCaseMock,
			createPaymentUseCase: createPaymentUseCaseMock,
		}

		order := domain.Order{}

		payment := result.PaymentProcessed{
			PaymentId:   uuid.New(),
			PaymentData: "mocked",
		}

		createPaymentUseCaseMock.On("CreatePayment", mock.Anything, mock.Anything).Return(&payment, nil)

		updateOrderUseCaseMock.On("Update", mock.Anything, payment.OrderId, valueobject.PaymentCreated, &payment.PaymentId).Return(errors.New("SOME_ERROR"))

		paymentCreated, err := processPaymentUseCase.ProcessPayment(context.TODO(), order)

		assert.NotNil(t, err)
		assert.Nil(t, paymentCreated)

		createPaymentUseCaseMock.AssertExpectations(t)
		createPaymentUseCaseMock.AssertCalled(t, "CreatePayment", mock.Anything, mock.Anything)

		updateOrderUseCaseMock.AssertExpectations(t)
		updateOrderUseCaseMock.AssertCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	})
}
