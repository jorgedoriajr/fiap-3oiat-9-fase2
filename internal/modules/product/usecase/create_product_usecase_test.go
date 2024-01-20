package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	ingredientDomain "hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/product/domain"
	"hamburgueria/internal/modules/product/usecase/command"
	mocks2 "hamburgueria/tests/mocks/modules/ingredient/ports/output"
	mocks "hamburgueria/tests/mocks/modules/product/ports/output"
	"testing"
)

func TestCreateProductUseCase(t *testing.T) {

	t.Run(`should create product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		ingredientPersistenceMock := mocks2.NewIngredientPersistencePort(t)
		createProductUseCase := CreateProductUseCase{
			productPersistencePort:    productPersistenceMock,
			productCategoryPort:       productCategoryPersistenceMock,
			ingredientPersistencePort: ingredientPersistenceMock,
		}

		ingredientTypeProductCategory := []ingredientDomain.IngredientTypeProductCategory{
			{
				Id:              uuid.New(),
				IngredientType:  "Type",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
		}

		ingredient := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		ingredient2 := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 2,
			Name:   "Ingredient2",
			Amount: 1500,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		productCategoryPersistenceMock.On("GetConfig", mock.Anything, "Category").Return(
			&domain.ProductCategory{
				Name:                    "Category",
				AcceptCustom:            true,
				ConfigByProductCategory: ingredientTypeProductCategory,
			}, nil)

		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)
		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 2).Return(&ingredient2, nil)

		productPersistenceMock.On("Create", mock.Anything, mock.Anything).Return(nil)

		productPersistenceMock.On("CheckProductExists", mock.Anything, mock.Anything).Return(nil, nil)

		productCreated, err := createProductUseCase.AddProduct(context.TODO(), command.CreateProductCommand{
			Name:        "Product",
			Description: "Description Product",
			Category:    "Category",
			Menu:        true,
			ImgPath:     "https://imgpath.com",
			Ingredients: []command.Ingredient{
				{
					Number:   1,
					Quantity: 1,
				},
				{
					Number:   2,
					Quantity: 1,
				},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, productCreated)

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetConfig", mock.Anything, "Category")

		ingredientPersistenceMock.AssertExpectations(t)
		ingredientPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		ingredientPersistenceMock.AssertNumberOfCalls(t, "GetByNumber", 2)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
		productPersistenceMock.AssertCalled(t, "CheckProductExists", mock.Anything, mock.Anything)
	})

	t.Run(`should not create product when doesn't have all ingredients'`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		ingredientPersistenceMock := mocks2.NewIngredientPersistencePort(t)
		createProductUseCase := CreateProductUseCase{
			productPersistencePort:    productPersistenceMock,
			productCategoryPort:       productCategoryPersistenceMock,
			ingredientPersistencePort: ingredientPersistenceMock,
		}

		ingredientTypeProductCategory := []ingredientDomain.IngredientTypeProductCategory{
			{
				Id:              uuid.New(),
				IngredientType:  "Type",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
		}

		mandatoryCategories := []ingredientDomain.IngredientTypeProductCategory{
			{
				Id:              uuid.New(),
				IngredientType:  "Type",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
			{
				Id:              uuid.New(),
				IngredientType:  "Type2",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
		}

		ingredient := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		ingredient2 := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 2,
			Name:   "Ingredient2",
			Amount: 1500,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		productCategoryPersistenceMock.On("GetConfig", mock.Anything, "Category").Return(&domain.ProductCategory{
			Name:                    "Category",
			AcceptCustom:            true,
			ConfigByProductCategory: mandatoryCategories,
		}, nil)

		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)
		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 2).Return(&ingredient2, nil)

		productCreated, err := createProductUseCase.AddProduct(context.TODO(), command.CreateProductCommand{
			Name:        "Product",
			Description: "Description Product",
			Category:    "Category",
			Menu:        true,
			ImgPath:     "https://imgpath.com",
			Ingredients: []command.Ingredient{
				{
					Number:   1,
					Quantity: 1,
				},
				{
					Number:   2,
					Quantity: 1,
				},
			},
		})

		assert.NotNil(t, err)
		assert.Nil(t, productCreated)

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetConfig", mock.Anything, "Category")

		ingredientPersistenceMock.AssertExpectations(t)
		ingredientPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		ingredientPersistenceMock.AssertNumberOfCalls(t, "GetByNumber", 2)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
		productPersistenceMock.AssertNotCalled(t, "CheckProductExists", mock.Anything, mock.Anything)
	})

	t.Run(`should not create product when quantity exceeds the limit`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		ingredientPersistenceMock := mocks2.NewIngredientPersistencePort(t)
		createProductUseCase := CreateProductUseCase{
			productPersistencePort:    productPersistenceMock,
			productCategoryPort:       productCategoryPersistenceMock,
			ingredientPersistencePort: ingredientPersistenceMock,
		}

		ingredientTypeProductCategory := []ingredientDomain.IngredientTypeProductCategory{
			{
				Id:              uuid.New(),
				IngredientType:  "Type",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
		}

		ingredient := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		ingredient2 := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 2,
			Name:   "Ingredient2",
			Amount: 1500,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		productCategoryPersistenceMock.On("GetConfig", mock.Anything, "Category").Return(
			&domain.ProductCategory{
				Name:                    "Category",
				AcceptCustom:            true,
				ConfigByProductCategory: ingredientTypeProductCategory,
			}, nil)

		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)
		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 2).Return(&ingredient2, nil)

		productCreated, err := createProductUseCase.AddProduct(context.TODO(), command.CreateProductCommand{
			Name:        "Product",
			Description: "Description Product",
			Category:    "Category",
			Menu:        true,
			ImgPath:     "https://imgpath.com",
			Ingredients: []command.Ingredient{
				{
					Number:   1,
					Quantity: 1,
				},
				{
					Number:   2,
					Quantity: 5,
				},
			},
		})

		assert.NotNil(t, err)
		assert.Nil(t, productCreated)

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetConfig", mock.Anything, "Category")

		ingredientPersistenceMock.AssertExpectations(t)
		ingredientPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		ingredientPersistenceMock.AssertNumberOfCalls(t, "GetByNumber", 2)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
		productPersistenceMock.AssertNotCalled(t, "CheckProductExists", mock.Anything, mock.Anything)
	})

	t.Run(`should not create product when ingredient is from another category`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		ingredientPersistenceMock := mocks2.NewIngredientPersistencePort(t)
		createProductUseCase := CreateProductUseCase{
			productPersistencePort:    productPersistenceMock,
			productCategoryPort:       productCategoryPersistenceMock,
			ingredientPersistencePort: ingredientPersistenceMock,
		}

		ingredientTypeProductCategory := []ingredientDomain.IngredientTypeProductCategory{
			{
				Id:              uuid.New(),
				IngredientType:  "Type",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
		}

		ingredient := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		ingredient2 := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 2,
			Name:   "Ingredient2",
			Amount: 1500,
			Type: ingredientDomain.IngredientType{
				Name: "Type",
				ConfigByProductCategory: []ingredientDomain.IngredientTypeProductCategory{
					{
						Id:              uuid.New(),
						IngredientType:  "Type2",
						Optional:        false,
						MaxQtd:          2,
						ProductCategory: "Category2",
					},
				},
			},
		}

		productCategoryPersistenceMock.On("GetConfig", mock.Anything, "Category").Return(
			&domain.ProductCategory{
				Name:                    "Category",
				AcceptCustom:            true,
				ConfigByProductCategory: ingredientTypeProductCategory,
			}, nil)

		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)
		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 2).Return(&ingredient2, nil)

		productCreated, err := createProductUseCase.AddProduct(context.TODO(), command.CreateProductCommand{
			Name:        "Product",
			Description: "Description Product",
			Category:    "Category",
			Menu:        true,
			ImgPath:     "https://imgpath.com",
			Ingredients: []command.Ingredient{
				{
					Number:   1,
					Quantity: 1,
				},
				{
					Number:   2,
					Quantity: 1,
				},
			},
		})

		assert.NotNil(t, err)
		assert.Nil(t, productCreated)

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetConfig", mock.Anything, "Category")

		ingredientPersistenceMock.AssertExpectations(t)
		ingredientPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		ingredientPersistenceMock.AssertNumberOfCalls(t, "GetByNumber", 2)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
		productPersistenceMock.AssertNotCalled(t, "CheckProductExists", mock.Anything, mock.Anything)
	})

	t.Run(`should return existent product when already exists`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		ingredientPersistenceMock := mocks2.NewIngredientPersistencePort(t)
		createProductUseCase := CreateProductUseCase{
			productPersistencePort:    productPersistenceMock,
			productCategoryPort:       productCategoryPersistenceMock,
			ingredientPersistencePort: ingredientPersistenceMock,
		}

		ingredientTypeProductCategory := []ingredientDomain.IngredientTypeProductCategory{
			{
				Id:              uuid.New(),
				IngredientType:  "Type",
				Optional:        false,
				MaxQtd:          2,
				ProductCategory: "Category",
			},
		}

		ingredient := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		ingredient2 := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 2,
			Name:   "Ingredient2",
			Amount: 1500,
			Type: ingredientDomain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: ingredientTypeProductCategory,
			},
		}

		productCategoryPersistenceMock.On("GetConfig", mock.Anything, "Category").Return(
			&domain.ProductCategory{
				Name:                    "Category",
				AcceptCustom:            true,
				ConfigByProductCategory: ingredientTypeProductCategory,
			}, nil)

		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)
		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 2).Return(&ingredient2, nil)

		productPersistenceMock.On("CheckProductExists", mock.Anything, mock.Anything).Return(&domain.Product{}, nil)

		productCreated, err := createProductUseCase.AddProduct(context.TODO(), command.CreateProductCommand{
			Name:        "Product",
			Description: "Description Product",
			Category:    "Category",
			Menu:        true,
			ImgPath:     "https://imgpath.com",
			Ingredients: []command.Ingredient{
				{
					Number:   1,
					Quantity: 1,
				},
				{
					Number:   2,
					Quantity: 1,
				},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, productCreated)

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetConfig", mock.Anything, "Category")

		ingredientPersistenceMock.AssertExpectations(t)
		ingredientPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		ingredientPersistenceMock.AssertNumberOfCalls(t, "GetByNumber", 2)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
		productPersistenceMock.AssertCalled(t, "CheckProductExists", mock.Anything, mock.Anything)
	})
}
