package controller

import (
	"gorm.io/gorm"
	paymentDatabase "hamburgueria/internal/modules/payment/infra/database"
	"hamburgueria/internal/modules/payment/port/input"
	paymentUseCase "hamburgueria/internal/modules/payment/usecase"
	"hamburgueria/pkg/logger"
	"sync"
)

type PaymentUseCaseController struct {
	CreatePaymentStatus input.CreatePaymentStatusPort
}

var (
	paymentUseCaseControllerInstance *PaymentUseCaseController
	paymentUseCaseControllerOnce     sync.Once
)

func GetPaymentUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *PaymentUseCaseController {
	paymentUseCaseControllerOnce.Do(func() {
		paymentStatusPersistenceGateway := paymentDatabase.GetPaymentStatusPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
		updateOrderUseCase := GetOrderUseCaseController(readWriteDB, readOnlyDB).UpdateOrderUseCase

		createPaymentStatus := paymentUseCase.GetCreatePaymentStatusUseCase(paymentStatusPersistenceGateway, updateOrderUseCase, logger.Get())

		paymentUseCaseControllerInstance = &PaymentUseCaseController{
			CreatePaymentStatus: createPaymentStatus,
		}
	})

	return paymentUseCaseControllerInstance

}
