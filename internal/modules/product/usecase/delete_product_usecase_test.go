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

func TestDeleteProductUseCase(t *testing.T) {

	t.Run(`should delete product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		deleteProductUseCase := DeleteProductUseCase{
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

		productInactive := product
		productInactive.Active = false

		productPersistenceMock.On("Update", mock.Anything, productInactive).Return(nil)

		err := deleteProductUseCase.Inactive(context.TODO(), 1)

		assert.Nil(t, err)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		productPersistenceMock.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when delete product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		deleteProductUseCase := DeleteProductUseCase{
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
		productPersistenceMock.On("Update", mock.Anything, mock.Anything).Return(errors.New("SOME_ERROR"))

		err := deleteProductUseCase.Inactive(context.TODO(), 1)

		assert.NotNil(t, err)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		productPersistenceMock.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when failed to found product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		deleteProductUseCase := DeleteProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(nil, errors.New("SOME_ERROR"))

		err := deleteProductUseCase.Inactive(context.TODO(), 1)

		assert.NotNil(t, err)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		productPersistenceMock.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
	})
}
