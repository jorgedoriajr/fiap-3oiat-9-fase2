package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/ingredient/domain"
	mocks "hamburgueria/tests/mocks/modules/ingredient/ports/output"
	"testing"
)

func TestFindIngredientUseCase(t *testing.T) {

	t.Run(`should find all ingredients`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredient := domain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: domain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: nil,
			},
		}

		ingredient2 := domain.Ingredient{
			ID:     uuid.New(),
			Number: 2,
			Name:   "Ingredient2",
			Amount: 1500,
			Type: domain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: nil,
			},
		}

		ingredientRepositoryMock.On("GetAll", mock.Anything).Return([]domain.Ingredient{ingredient, ingredient2}, nil)

		ingredients, err := findIngredientUseCase.FindAllIngredients(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(ingredients), 2)
		assert.Equal(t, ingredients[0].Name, "Ingredient")
		assert.Equal(t, ingredients[1].Name, "Ingredient2")
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should return error when GetAll failed`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredientRepositoryMock.On("GetAll", mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		ingredients, err := findIngredientUseCase.FindAllIngredients(context.TODO())

		assert.Nil(t, ingredients)
		assert.NotNil(t, err)
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should return empty`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredientRepositoryMock.On("GetAll", mock.Anything).Return([]domain.Ingredient{}, nil)

		ingredients, err := findIngredientUseCase.FindAllIngredients(context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, len(ingredients), 0)
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetAll", mock.Anything)
	})

	t.Run(`should find ingredient by type`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredient := domain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: domain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: nil,
			},
		}

		ingredientRepositoryMock.On("GetByType", mock.Anything, "Type").Return([]domain.Ingredient{ingredient}, nil)

		ingredientFound, err := findIngredientUseCase.FindIngredientByType(context.TODO(), "Type")

		assert.Nil(t, err)
		assert.Equal(t, len(ingredientFound), 1)
		assert.Equal(t, ingredientFound[0].Name, "Ingredient")
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetByType", mock.Anything, mock.Anything)
	})

	t.Run(`should return when when find ingredient by type with error`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredientRepositoryMock.On("GetByType", mock.Anything, "Type").Return(nil, errors.New("SOME_ERROR"))

		ingredientFound, err := findIngredientUseCase.FindIngredientByType(context.TODO(), "Type")

		assert.Nil(t, ingredientFound)
		assert.NotNil(t, err)
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetByType", mock.Anything, mock.Anything)
	})

	t.Run(`should find ingredient by number`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredient := domain.Ingredient{
			ID:     uuid.New(),
			Number: 1,
			Name:   "Ingredient",
			Amount: 1000,
			Type: domain.IngredientType{
				Name:                    "Type",
				ConfigByProductCategory: nil,
			},
		}

		ingredientRepositoryMock.On("GetByNumber", mock.Anything, 1).Return(&ingredient, nil)

		ingredientFound, err := findIngredientUseCase.FindIngredientByNumber(context.TODO(), 1)

		assert.Nil(t, err)
		assert.Equal(t, ingredientFound.Name, "Ingredient")
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
	})

	t.Run(`should return error when find ingredient by number with error`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		findIngredientUseCase := FindIngredientUseCase{
			ingredientPersistence: ingredientRepositoryMock,
		}

		ingredientRepositoryMock.On("GetByNumber", mock.Anything, 1).Return(nil, errors.New("SOME_ERROR"))

		ingredientFound, err := findIngredientUseCase.FindIngredientByNumber(context.TODO(), 1)

		assert.NotNil(t, err)
		assert.Nil(t, ingredientFound)
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "GetByNumber", mock.Anything, mock.Anything)
	})
}
