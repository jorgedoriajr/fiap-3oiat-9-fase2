package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	ingredientDomain "hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/product/domain"
	"hamburgueria/tests/mocks"
	"testing"
	"time"
)

func TestFindProductUseCase(t *testing.T) {

	t.Run(`should find all products`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productId := uuid.New()
		product := domain.Product{
			ID:          productId,
			Number:      1,
			Name:        "Product",
			Amount:      4000,
			Description: "Product Description",
			Category:    domain.ProductCategory{Name: "Category"},
			Menu:        true,
			ImgPath:     "https://image.com",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Ingredients: []domain.ProductIngredient{
				{
					ID:        uuid.New(),
					ProductId: productId,
					Ingredient: ingredientDomain.Ingredient{
						ID:     uuid.New(),
						Number: 1,
						Name:   "Ingredient",
						Amount: 1000,
						Type:   ingredientDomain.IngredientType{Name: "Type"},
					},
					Quantity: 4,
					Amount:   4000,
				},
			},
			Active: true,
		}

		productPersistenceMock.On("GetAll", mock.Anything).Return([]domain.Product{product}, nil)

		products, err := listProductUseCase.FindAllProducts(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(products), 1)
		assert.Equal(t, products[0].Name, "Product")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should return error`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productPersistenceMock.On("GetAll", mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		products, err := listProductUseCase.FindAllProducts(context.TODO())

		assert.NotNil(t, err)
		assert.Nil(t, products)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should find products by category`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productId := uuid.New()
		product := domain.Product{
			ID:          productId,
			Number:      1,
			Name:        "Product",
			Amount:      4000,
			Description: "Product Description",
			Category:    domain.ProductCategory{Name: "Category"},
			Menu:        true,
			ImgPath:     "https://image.com",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Ingredients: []domain.ProductIngredient{
				{
					ID:        uuid.New(),
					ProductId: productId,
					Ingredient: ingredientDomain.Ingredient{
						ID:     uuid.New(),
						Number: 1,
						Name:   "Ingredient",
						Amount: 1000,
						Type:   ingredientDomain.IngredientType{Name: "Type"},
					},
					Quantity: 4,
					Amount:   4000,
				},
			},
			Active: true,
		}

		productPersistenceMock.On("GetByCategory", mock.Anything, "Category").Return([]domain.Product{product}, nil)

		products, err := listProductUseCase.FindByCategory(context.TODO(), "Category")

		assert.Nil(t, err)
		assert.Equal(t, len(products), 1)
		assert.Equal(t, products[0].Name, "Product")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByCategory", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when find by category`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productPersistenceMock.On("GetByCategory", mock.Anything, "Category").Return(nil, errors.New("SOME_ERROR"))

		products, err := listProductUseCase.FindByCategory(context.TODO(), "Category")

		assert.NotNil(t, err)
		assert.Nil(t, products)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByCategory", mock.Anything, mock.Anything)
	})

	t.Run(`should find product by number`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productId := uuid.New()
		product := domain.Product{
			ID:          productId,
			Number:      1,
			Name:        "Product",
			Amount:      4000,
			Description: "Product Description",
			Category:    domain.ProductCategory{Name: "Category"},
			Menu:        true,
			ImgPath:     "https://image.com",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Ingredients: []domain.ProductIngredient{
				{
					ID:        uuid.New(),
					ProductId: productId,
					Ingredient: ingredientDomain.Ingredient{
						ID:     uuid.New(),
						Number: 1,
						Name:   "Ingredient",
						Amount: 1000,
						Type:   ingredientDomain.IngredientType{Name: "Type"},
					},
					Quantity: 4,
					Amount:   4000,
				},
			},
			Active: true,
		}

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&product, nil)

		productFound, err := listProductUseCase.FindByNumber(context.TODO(), 1)

		assert.Nil(t, err)
		assert.Equal(t, productFound.Name, "Product")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when find by number`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(nil, errors.New("SOME_ERROR"))

		products, err := listProductUseCase.FindByNumber(context.TODO(), 1)

		assert.NotNil(t, err)
		assert.Nil(t, products)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
	})

	t.Run(`should find product by ID`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productId := uuid.New()
		product := domain.Product{
			ID:          productId,
			Number:      1,
			Name:        "Product",
			Amount:      4000,
			Description: "Product Description",
			Category:    domain.ProductCategory{Name: "Category"},
			Menu:        true,
			ImgPath:     "https://image.com",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Ingredients: []domain.ProductIngredient{
				{
					ID:        uuid.New(),
					ProductId: productId,
					Ingredient: ingredientDomain.Ingredient{
						ID:     uuid.New(),
						Number: 1,
						Name:   "Ingredient",
						Amount: 1000,
						Type:   ingredientDomain.IngredientType{Name: "Type"},
					},
					Quantity: 4,
					Amount:   4000,
				},
			},
			Active: true,
		}

		productPersistenceMock.On("GetByID", mock.Anything, productId).Return(&product, nil)

		productFound, err := listProductUseCase.FindByID(context.TODO(), productId)

		assert.Nil(t, err)
		assert.Equal(t, productFound.Name, "Product")

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByID", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when find by id`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		listProductUseCase := FindProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productPersistenceMock.On("GetByID", mock.Anything, mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		products, err := listProductUseCase.FindByID(context.TODO(), uuid.New())

		assert.NotNil(t, err)
		assert.Nil(t, products)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByID", mock.Anything, mock.Anything)
	})
}
