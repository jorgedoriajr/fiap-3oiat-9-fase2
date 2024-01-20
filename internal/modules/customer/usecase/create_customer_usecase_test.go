package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/customer/domain"
	"hamburgueria/internal/modules/customer/usecase/command"
	mocks "hamburgueria/tests/mocks/modules/customer/port/output"

	"testing"
)

func TestCreateCustomerUseCase(t *testing.T) {

	t.Run(`should add customer`, func(t *testing.T) {
		repositoryMock := mocks.NewCustomerPersistencePort(t)
		allowListService := CreateCustomerUseCase{
			customerPersistence: repositoryMock,
		}

		customer := command.CreateCustomerCommand{
			Document: "58642725826",
			Name:     "Name",
			Phone:    "11999999999",
		}

		repositoryMock.On("Create", mock.Anything, mock.MatchedBy(func(c domain.Customer) bool {
			return c.Document == "58642725826" &&
				c.Name == "Name" &&
				c.Phone == "11999999999"
		})).Return(nil)

		err := allowListService.AddCustomer(context.TODO(), customer)

		assert.Nil(t, err)
		repositoryMock.AssertExpectations(t)
		repositoryMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run(`should not add customer if error`, func(t *testing.T) {
		repositoryMock := mocks.NewCustomerPersistencePort(t)
		allowListService := CreateCustomerUseCase{
			customerPersistence: repositoryMock,
		}

		customer := command.CreateCustomerCommand{
			Document: "58642725826",
			Name:     "Name",
			Phone:    "11999999999",
		}

		repositoryMock.On("Create", mock.Anything, mock.Anything).Return(errors.New("SOME_ERROR"))

		err := allowListService.AddCustomer(context.TODO(), customer)

		assert.NotNil(t, err)
		repositoryMock.AssertExpectations(t)
		repositoryMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})

}
