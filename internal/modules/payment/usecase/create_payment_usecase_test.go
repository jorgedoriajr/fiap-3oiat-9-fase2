package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hamburgueria/internal/modules/payment/usecase/command"
	"testing"
)

func TestCreatePaymentUseCase(t *testing.T) {

	t.Run(`should create payment`, func(t *testing.T) {
		createPaymentUseCase := CreatePaymentUseCase{}
		paymentCreated, err := createPaymentUseCase.CreatePayment(context.TODO(), command.CreatePaymentCommand{})
		assert.Nil(t, err)
		assert.Equal(t, "mocked", paymentCreated.PaymentData)
	})
}
