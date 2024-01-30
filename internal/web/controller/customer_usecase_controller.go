package controller

import (
	"gorm.io/gorm"
	"hamburgueria/internal/modules/customer/infra/database"
	"hamburgueria/internal/modules/customer/port/input"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	"hamburgueria/pkg/logger"
	"sync"
)

type CustomerUseCaseController struct {
	CreateCustomerUseCase input.CreateCustomerPort
	GetCustomerUseCase    input.GetCustomerPort
}

var (
	customerUseCaseControllerInstance *CustomerUseCaseController
	customerUseCaseControllerOnce     sync.Once
)

func GetCustomerUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *CustomerUseCaseController {
	customerUseCaseControllerOnce.Do(func() {
		customerPersistence := database.GetCustomerPersistence(readWriteDB, readOnlyDB, logger.Get())

		customerUseCaseControllerInstance = &CustomerUseCaseController{
			CreateCustomerUseCase: customerUseCase.GetCreateCustomerUseCase(customerPersistence),
			GetCustomerUseCase:    customerUseCase.GetGetCustomerUseCase(customerPersistence),
		}
	})

	return customerUseCaseControllerInstance
}
