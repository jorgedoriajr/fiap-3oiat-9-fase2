package controller

import (
	"gorm.io/gorm"
	"hamburgueria/internal/modules/customer/infra/database"
	orderDatabase "hamburgueria/internal/modules/order/infra/database"
	"hamburgueria/internal/modules/order/usecase"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago"
	paymentDatabase "hamburgueria/internal/modules/payment/infra/database"
	paymentUseCase "hamburgueria/internal/modules/payment/usecase"
	productDatabase "hamburgueria/internal/modules/product/infra/database"
	"hamburgueria/pkg/httpclient"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/starter"
)

type OrderUseCaseController struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *OrderUseCaseController {
	orderPersistence := orderDatabase.GetOrderPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	productPersistence := productDatabase.GetProductPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	customerPersistence := database.GetCustomerPersistence(readWriteDB, readOnlyDB, logger.Get())
	updateOrderUseCase := usecase.GetUpdateOrderUseCase(orderPersistence)

	paymentPersistance := paymentDatabase.GetPaymentPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

	mercadoPagoClient := mercadopago.GetCreateMercadoPagoClient(
		httpclient.GetClient("mercadoPago"),
		starter.GetConfigRoot().MercadoPago,
		logger.Get(),
	)

	createPaymentUseCase := paymentUseCase.GetCreatePaymentUseCase(mercadoPagoClient, &paymentPersistance)
	processPaymentUseCase := usecase.GetProcessPaymentUseCase(updateOrderUseCase, createPaymentUseCase)

	createOrderUseCase := usecase.GetCreateOrderUseCase(
		productPersistence,
		orderPersistence,
		processPaymentUseCase,
		customerPersistence,
	)

	listOrderUseCase := usecase.GetListOrderUseCase(orderPersistence)

	return &OrderUseCaseController{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}

}
