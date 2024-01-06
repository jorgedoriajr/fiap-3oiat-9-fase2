package usecase

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/usecase/command"
	"hamburgueria/internal/modules/payment/usecase/result"
	"sync"
)

type CreatePaymentUseCase struct {
}

func (p CreatePaymentUseCase) CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	return &result.PaymentProcessed{
		PaymentId:   uuid.New(),
		PaymentData: "mocked",
	}, nil
}

var (
	processPaymentUseCase     CreatePaymentUseCase
	processPaymentUseCaseOnce sync.Once
)

func GetCreatePaymentUseCase() input.CreatePaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCase = CreatePaymentUseCase{}
	})
	return processPaymentUseCase
}
