package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/product/domain"
	mocks "hamburgueria/tests/mocks/modules/product/ports/output"
	"testing"
)

func TestGetProductCategoryUseCase(t *testing.T) {

	t.Run(`should find all categories`, func(t *testing.T) {
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		getProductCategoryUseCase := GetProductCategoryUseCase{
			productCategoryPersistenceGateway: productCategoryPersistenceMock,
		}

		productCategory := domain.ProductCategory{
			Name:                    "Category",
			AcceptCustom:            false,
			ConfigByProductCategory: nil,
		}

		productCategory2 := domain.ProductCategory{
			Name:                    "Category2",
			AcceptCustom:            true,
			ConfigByProductCategory: nil,
		}

		productCategoryPersistenceMock.On("GetAll", mock.Anything).Return([]domain.ProductCategory{productCategory, productCategory2}, nil)

		categories, err := getProductCategoryUseCase.FindAll(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(categories), 2)
		assert.Equal(t, categories[0].Name, "Category")
		assert.Equal(t, categories[1].Name, "Category2")

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should return error`, func(t *testing.T) {
		productCategoryPersistenceMock := mocks.NewProductCategoryPersistencePort(t)
		getProductCategoryUseCase := GetProductCategoryUseCase{
			productCategoryPersistenceGateway: productCategoryPersistenceMock,
		}

		productCategoryPersistenceMock.On("GetAll", mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		categories, err := getProductCategoryUseCase.FindAll(context.TODO())

		assert.Nil(t, categories)
		assert.NotNil(t, err)

		productCategoryPersistenceMock.AssertExpectations(t)
		productCategoryPersistenceMock.AssertCalled(t, "GetAll", mock.Anything)
	})
}
