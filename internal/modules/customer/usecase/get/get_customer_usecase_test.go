package get

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/tests/mocks"
	"testing"
	"time"
)

func TestGetCustomerUseCase(t *testing.T) {

	t.Run(`should get customer`, func(t *testing.T) {
		repositoryMock := mocks.NewCustomerPersistencePort(t)
		allowListService := GetCustomerUseCase{
			CustomerPersistence: repositoryMock,
		}

		document := "58642725826"

		response := entity.Customer{
			Document:  "58642725826",
			Name:      "Name",
			Phone:     "11999999999",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		repositoryMock.On("Get", mock.Anything, document).Return(&response, nil)

		customer, err := allowListService.GetCustomer(context.TODO(), document)

		assert.Nil(t, err)
		assert.Equal(t, customer.Document, response.Document)
		assert.Equal(t, customer.Phone, response.Phone)
		assert.Equal(t, customer.Name, response.Name)
		repositoryMock.AssertExpectations(t)
		repositoryMock.AssertCalled(t, "Get", mock.Anything, document)
	})

	t.Run(`should return nil when not found`, func(t *testing.T) {
		repositoryMock := mocks.NewCustomerPersistencePort(t)
		allowListService := GetCustomerUseCase{
			CustomerPersistence: repositoryMock,
		}

		document := "58642725826"

		repositoryMock.On("Get", mock.Anything, document).Return(nil, nil)

		customer, err := allowListService.GetCustomer(context.TODO(), document)

		assert.Nil(t, err)
		assert.Nil(t, customer)
		repositoryMock.AssertExpectations(t)
		repositoryMock.AssertCalled(t, "Get", mock.Anything, document)
	})

	t.Run(`should return error if something wrong while trying to get`, func(t *testing.T) {
		repositoryMock := mocks.NewCustomerPersistencePort(t)
		allowListService := GetCustomerUseCase{
			CustomerPersistence: repositoryMock,
		}

		document := "58642725826"

		repositoryMock.On("Get", mock.Anything, document).Return(nil, errors.New("SOME_ERROR"))

		customer, err := allowListService.GetCustomer(context.TODO(), document)

		assert.NotNil(t, err)
		assert.Nil(t, customer)
		repositoryMock.AssertExpectations(t)
		repositoryMock.AssertCalled(t, "Get", mock.Anything, document)
	})

}
