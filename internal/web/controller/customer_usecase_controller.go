package controller

import (
	"gorm.io/gorm"
	"hamburgueria/internal/modules/customer/infra/database"
	"hamburgueria/internal/modules/customer/port/input"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	"hamburgueria/pkg/logger"
)

type CustomerUseCaseController struct {
	CreateCustomerUseCase input.CreateCustomerPort
	GetCustomerUseCase    input.GetCustomerPort
}

func NewCustomerUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *CustomerUseCaseController {
	customerPersistence := database.GetCustomerPersistence(readWriteDB, readOnlyDB, logger.Get())

	return &CustomerUseCaseController{
		CreateCustomerUseCase: customerUseCase.GetCreateCustomerUseCase(customerPersistence),
		GetCustomerUseCase:    customerUseCase.GetGetCustomerUseCase(customerPersistence),
	}

}
