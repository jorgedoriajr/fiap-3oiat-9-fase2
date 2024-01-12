package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/ingredient/usecase/command"
	mocks "hamburgueria/tests/mocks/modules/ingredient/ports/output"
	"testing"
)

func TestCreateIngredientUseCase(t *testing.T) {

	t.Run(`should add ingredient`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		ingredientTypeRepositoryMock := mocks.NewIngredientTypePersistencePort(t)
		createIngredientUseCase := CreateIngredientUseCase{
			ingredientPersistence:     ingredientRepositoryMock,
			ingredientTypePersistence: ingredientTypeRepositoryMock,
		}

		ingredientCommand := command.CreateIngredientCommand{
			Name:   "Ingredient",
			Amount: 1000,
			Type:   "Type",
		}

		ingredientType := domain.IngredientType{
			Name:                    "Type",
			ConfigByProductCategory: []domain.IngredientTypeProductCategory{},
		}

		ingredientTypeRepositoryMock.On("GetByName", mock.Anything, "Type").Return(&ingredientType, nil)

		ingredientRepositoryMock.On("Create", mock.Anything, mock.MatchedBy(func(c domain.Ingredient) bool {
			return c.Amount == 1000 &&
				c.Name == "Ingredient" &&
				c.Type.Name == "Type"
		})).Return(nil)

		ingredientCreated, err := createIngredientUseCase.AddIngredient(context.TODO(), ingredientCommand)

		assert.Nil(t, err)
		assert.Equal(t, ingredientCreated.Name, "Ingredient")
		assert.Equal(t, ingredientCreated.Amount, 1000)
		assert.Equal(t, ingredientCreated.Type, "Type")
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)

		ingredientTypeRepositoryMock.AssertExpectations(t)
		ingredientTypeRepositoryMock.AssertCalled(t, "GetByName", mock.Anything, mock.Anything)
	})

	t.Run(`should not add ingredient if error`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		ingredientTypeRepositoryMock := mocks.NewIngredientTypePersistencePort(t)
		createIngredientUseCase := CreateIngredientUseCase{
			ingredientPersistence:     ingredientRepositoryMock,
			ingredientTypePersistence: ingredientTypeRepositoryMock,
		}

		ingredientCommand := command.CreateIngredientCommand{
			Name:   "Ingredient",
			Amount: 1000,
			Type:   "Type",
		}

		ingredientType := domain.IngredientType{
			Name:                    "Type",
			ConfigByProductCategory: []domain.IngredientTypeProductCategory{},
		}

		ingredientTypeRepositoryMock.On("GetByName", mock.Anything, "Type").Return(&ingredientType, nil)

		ingredientRepositoryMock.On("Create", mock.Anything, mock.MatchedBy(func(c domain.Ingredient) bool {
			return c.Amount == 1000 &&
				c.Name == "Ingredient" &&
				c.Type.Name == "Type"
		})).Return(errors.New("SOME_ERROR"))

		ingredientCreated, err := createIngredientUseCase.AddIngredient(context.TODO(), ingredientCommand)

		assert.NotNil(t, err)
		assert.Nil(t, ingredientCreated)
		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)

		ingredientTypeRepositoryMock.AssertExpectations(t)
		ingredientTypeRepositoryMock.AssertCalled(t, "GetByName", mock.Anything, mock.Anything)
	})

	t.Run(`should not add ingredient if error to get type`, func(t *testing.T) {
		ingredientRepositoryMock := mocks.NewIngredientPersistencePort(t)
		ingredientTypeRepositoryMock := mocks.NewIngredientTypePersistencePort(t)
		createIngredientUseCase := CreateIngredientUseCase{
			ingredientPersistence:     ingredientRepositoryMock,
			ingredientTypePersistence: ingredientTypeRepositoryMock,
		}

		ingredientCommand := command.CreateIngredientCommand{
			Name:   "Ingredient",
			Amount: 1000,
			Type:   "Type",
		}

		ingredientTypeRepositoryMock.On("GetByName", mock.Anything, mock.Anything).Return(nil, errors.New("SOME_ERROR"))

		result, err := createIngredientUseCase.AddIngredient(context.TODO(), ingredientCommand)

		assert.NotNil(t, err)
		assert.Nil(t, result)
		ingredientTypeRepositoryMock.AssertExpectations(t)
		ingredientTypeRepositoryMock.AssertCalled(t, "GetByName", mock.Anything, mock.Anything)

		ingredientRepositoryMock.AssertExpectations(t)
		ingredientRepositoryMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

}
