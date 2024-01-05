package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	ingredientDomain "hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/product/domain"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/tests/mocks"
	"testing"
	"time"
)

func TestUpdateProductUseCase(t *testing.T) {

	t.Run(`should update product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		ingredientPersistenceMock := mocks.NewIngredientPersistencePort(t)
		updateProductUseCase := UpdateProductUseCase{
			productPersistencePort:    productPersistenceMock,
			ingredientPersistencePort: ingredientPersistenceMock,
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

		edited := "Edited"

		productCommand := command.UpdateProductCommand{
			Number:      1,
			Name:        &edited,
			Description: &edited,
			Category:    nil,
			Menu:        nil,
			ImgPath:     nil,
			Ingredients: []command.Ingredient{
				{
					Number:   1,
					Quantity: 5,
				},
			},
		}

		ingredient := ingredientDomain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type:   ingredientDomain.IngredientType{Name: "Type"},
		}

		productPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&product, nil)
		ingredientPersistenceMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)

		productPersistenceMock.On("Update", mock.Anything, mock.MatchedBy(func(c domain.Product) bool {
			return c.Name == "Edited" &&
				c.Description == "Edited"
		})).Return(nil)

		err := updateProductUseCase.UpdateProduct(context.TODO(), productCommand)

		assert.Nil(t, err)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		productPersistenceMock.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when update product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		updateProductUseCase := UpdateProductUseCase{
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
			Ingredients: nil,
		}

		productPersistenceMock.On("GetByNumber", mock.Anything, 0).Return(&product, nil)
		productPersistenceMock.On("Update", mock.Anything, mock.Anything).Return(errors.New("SOME_ERROR"))

		err := updateProductUseCase.UpdateProduct(context.TODO(), command.UpdateProductCommand{})

		assert.NotNil(t, err)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		productPersistenceMock.AssertCalled(t, "Update", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when failed to found product`, func(t *testing.T) {
		productPersistenceMock := mocks.NewProductPersistencePort(t)
		updateProductUseCase := UpdateProductUseCase{
			productPersistencePort: productPersistenceMock,
		}

		productPersistenceMock.On("GetByNumber", mock.Anything, 0).Return(nil, errors.New("SOME_ERROR"))

		err := updateProductUseCase.UpdateProduct(context.TODO(), command.UpdateProductCommand{})

		assert.NotNil(t, err)

		productPersistenceMock.AssertExpectations(t)
		productPersistenceMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
		productPersistenceMock.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
	})
}
