package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/ingredient/domain"
	mocks "hamburgueria/tests/mocks/modules/ingredient/ports/output"

	"testing"
)

func TestFindIngredientTypeUseCase(t *testing.T) {

	t.Run(`should find all ingredient type`, func(t *testing.T) {
		ingredientTypeRepositoryMock := mocks.NewIngredientTypePersistencePort(t)
		findIngredientTypeUseCase := FindIngredientTypeUseCase{
			ingredientTypePersistenceGateway: ingredientTypeRepositoryMock,
		}

		type1 := domain.IngredientType{
			Name:                    "Type",
			ConfigByProductCategory: nil,
		}

		type2 := domain.IngredientType{
			Name:                    "Type2",
			ConfigByProductCategory: nil,
		}

		ingredientTypeRepositoryMock.On("GetAll", mock.Anything).Return([]domain.IngredientType{type1, type2}, nil)

		ingredientTypes, err := findIngredientTypeUseCase.FindAll(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(ingredientTypes), 2)
		assert.Equal(t, ingredientTypes[0].Name, "Type")
		assert.Equal(t, ingredientTypes[1].Name, "Type2")
		ingredientTypeRepositoryMock.AssertExpectations(t)
		ingredientTypeRepositoryMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should return error when GetAll failed`, func(t *testing.T) {
		ingredientTypeRepositoryMock := mocks.NewIngredientTypePersistencePort(t)
		findIngredientTypeUseCase := FindIngredientTypeUseCase{
			ingredientTypePersistenceGateway: ingredientTypeRepositoryMock,
		}

		ingredientTypeRepositoryMock.On("GetAll", mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		ingredientTypes, err := findIngredientTypeUseCase.FindAll(context.TODO())

		assert.Nil(t, ingredientTypes)
		assert.NotNil(t, err)
		ingredientTypeRepositoryMock.AssertExpectations(t)
		ingredientTypeRepositoryMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should return empty`, func(t *testing.T) {
		ingredientTypeRepositoryMock := mocks.NewIngredientTypePersistencePort(t)
		findIngredientTypeUseCase := FindIngredientTypeUseCase{
			ingredientTypePersistenceGateway: ingredientTypeRepositoryMock,
		}

		ingredientTypeRepositoryMock.On("GetAll", mock.Anything).Return([]domain.IngredientType{}, nil)

		ingredientTypes, err := findIngredientTypeUseCase.FindAll(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(ingredientTypes), 0)
		ingredientTypeRepositoryMock.AssertExpectations(t)
		ingredientTypeRepositoryMock.AssertCalled(t, "GetAll", mock.Anything)
	})
}
