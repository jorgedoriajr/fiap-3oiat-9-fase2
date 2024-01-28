package controller

import (
	"gorm.io/gorm"
	ingredientDatabase "hamburgueria/internal/modules/ingredient/infra/database"
	productDatabase "hamburgueria/internal/modules/product/infra/database"
	"hamburgueria/internal/modules/product/ports/input"
	"hamburgueria/internal/modules/product/usecase"
	"hamburgueria/pkg/logger"
)

var ()

type ProductUseCaseController struct {
	CreateProductUseCase      input.CreateProductUseCasePort
	DeleteProductUseCase      input.DeleteProductUseCasePort
	GetProductCategoryUseCase input.GetProductCategoryUseCasePort
	FindProductUseCase        input.FindProductUseCasePort
	UpdateProductUseCase      input.UpdateProductUseCasePort
}

func NewProductUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *ProductUseCaseController {
	productPersistence := productDatabase.GetProductPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	productCategoryPersistence := productDatabase.GetProductCategoryRepository(readWriteDB, readOnlyDB, logger.Get())
	ingredientPersistence := ingredientDatabase.GetIngredientPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

	findProductCategoryUseCase := usecase.NewGetProductCategoryUseCase(productCategoryPersistence)
	createProductUseCase := usecase.GetCreateProductUseCase(productPersistence, ingredientPersistence, productCategoryPersistence)
	deleteProductUseCase := usecase.GetDeleteProductUseCase(productPersistence)
	updateProductUseCase := usecase.GetUpdateProductUseCase(productPersistence, ingredientPersistence)
	findProductUseCase := usecase.NewFindProductUseCase(productPersistence)

	return &ProductUseCaseController{
		CreateProductUseCase:      createProductUseCase,
		DeleteProductUseCase:      deleteProductUseCase,
		GetProductCategoryUseCase: findProductCategoryUseCase,
		FindProductUseCase:        findProductUseCase,
		UpdateProductUseCase:      updateProductUseCase,
	}
}
