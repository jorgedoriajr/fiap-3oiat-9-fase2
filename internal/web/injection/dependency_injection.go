package injection

import (
	"hamburgueria/internal/modules/customer/infra/database"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	ingredientDatabase "hamburgueria/internal/modules/ingredient/infra/database"
	create2 "hamburgueria/internal/modules/ingredient/usecase"
	orderDatabase "hamburgueria/internal/modules/order/infra/database"
	orderUsecase "hamburgueria/internal/modules/order/usecase"
	"hamburgueria/internal/modules/payment/infra/client/mercadopago"
	paymentDatabase "hamburgueria/internal/modules/payment/infra/database"
	paymentUseCase "hamburgueria/internal/modules/payment/usecase"
	productDatabase "hamburgueria/internal/modules/product/infra/database"
	"hamburgueria/internal/modules/product/usecase"
	"hamburgueria/internal/web/api/rest/v1/customer"
	"hamburgueria/internal/web/api/rest/v1/ingredient"
	"hamburgueria/internal/web/api/rest/v1/ingredienttype"
	"hamburgueria/internal/web/api/rest/v1/order"
	"hamburgueria/internal/web/api/rest/v1/payment"
	"hamburgueria/internal/web/api/rest/v1/product"
	"hamburgueria/internal/web/api/rest/v1/productcategory"
	"hamburgueria/internal/web/api/swagger"
	"hamburgueria/pkg/httpclient"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
	"hamburgueria/pkg/starter"
)

type DependencyInjection struct {
	CustomerApi           *customer.Api
	ProductApi            *product.Api
	IngredientApi         *ingredient.Api
	OrderApi              *order.Api
	IngredientTypeApi     *ingredienttype.Api
	ProductCategoryApi    *productcategory.Api
	PaymentsStatusWebhook *payment.Webhook
	Swagger               *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	readWriteDB, readOnlyDB := sql.GetClient("readWrite"), sql.GetClient("readOnly")

	customerPersistence := database.GetCustomerPersistence(readWriteDB, readOnlyDB, logger.Get())

	ingredientTypePersistence := ingredientDatabase.GetIngredientTypePersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	ingredientPersistence := ingredientDatabase.GetIngredientPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

	productCategoryPersistence := productDatabase.GetProductCategoryRepository(readWriteDB, readOnlyDB, logger.Get())

	productPersistence := productDatabase.GetProductPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	findProductCategoryUseCase := usecase.NewGetProductCategoryUseCase(productCategoryPersistence)
	createProductUseCase := usecase.GetCreateProductUseCase(productPersistence, ingredientPersistence, productCategoryPersistence)
	deleteProductUseCase := usecase.GetDeleteProductUseCase(productPersistence)
	updateProductUseCase := usecase.GetUpdateProductUseCase(productPersistence, ingredientPersistence)
	findProductUseCase := usecase.NewFindProductUseCase(productPersistence)

	orderPersistence := orderDatabase.GetOrderPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

	paymentPersistance := paymentDatabase.GetPaymentPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	paymentStatusPersistance := paymentDatabase.GetPaymentStatusPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

	createPaymentStatusUseCase := paymentUseCase.GetCreatePaymentStatusUseCase(&paymentStatusPersistance)

	mercadoPagoClient := mercadopago.GetCreateMercadoPagoClient(
		httpclient.GetClient("mercadoPago"),
		starter.GetConfigRoot().MercadoPago,
		logger.Get(),
	)

	updateOrderUseCase := orderUsecase.GetUpdateOrderUseCase(orderPersistence)

	createPaymentUseCase := paymentUseCase.GetCreatePaymentUseCase(mercadoPagoClient, &paymentPersistance)
	processPaymentUseCase := orderUsecase.GetProcessPaymentUseCase(updateOrderUseCase, createPaymentUseCase)

	createOrderUseCase := orderUsecase.GetCreateOrderUseCase(
		productPersistence,
		orderPersistence,
		processPaymentUseCase,
		customerPersistence,
	)

	return DependencyInjection{
		CustomerApi: &customer.Api{
			CreateCustomerUseCase: customerUseCase.GetCreateCustomerUseCase(customerPersistence),
			GetCustomerUseCase:    customerUseCase.GetGetCustomerUseCase(customerPersistence),
		},
		ProductApi: &product.Api{
			CreateProductUseCase:  createProductUseCase,
			FindProductUseCase:    findProductUseCase,
			DeleteProductUseCase:  deleteProductUseCase,
			UpdatedProductUseCase: updateProductUseCase,
		},
		OrderApi: &order.Api{
			CreateOrderUseCase: createOrderUseCase,
			ListOrderUseCase:   orderUsecase.GetListOrderUseCase(orderPersistence),
		},
		IngredientApi: &ingredient.Api{
			CreateIngredientUseCase: create2.NewCreateIngredientUseCase(ingredientPersistence, ingredientTypePersistence),
			FindIngredientUseCase:   create2.NewFindIngredientUseCase(ingredientPersistence),
		},
		ProductCategoryApi: &productcategory.Api{GetProductCategoryUseCase: findProductCategoryUseCase},
		IngredientTypeApi: &ingredienttype.Api{
			FindIngredientTypeUseCase: create2.GetIngredientTypeUseCase(ingredientTypePersistence),
		},
		PaymentsStatusWebhook: &payment.Webhook{CreatePaymentStatusUseCase: createPaymentStatusUseCase},
		Swagger:               &swagger.Swagger{},
	}
}
