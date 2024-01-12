package customer

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hamburgueria/internal/modules/customer/usecase/result"
	mocks "hamburgueria/tests/mocks/modules/customer/port/input"

	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCustomerController(t *testing.T) {
	t.Run(`should get customer`, func(t *testing.T) {

		createCustomerPortMock := mocks.NewCreateCustomerPort(t)
		getCustomerPortMock := mocks.NewGetCustomerPort(t)
		controller := Controller{
			CreateCustomerUseCase: createCustomerPortMock,
			GetCustomerUseCase:    getCustomerPortMock,
		}

		document := "98932673667"
		customer := result.CustomerCreated{
			Document: document,
			Name:     "Name",
			Phone:    "11999999999",
		}

		echoServer := echo.New()
		testUrl := fmt.Sprintf("http://localhost:8080/v1/customers/%s", document)

		getCustomerPortMock.On("GetCustomer", mock.Anything, document).Return(&customer, nil)

		req, _ := http.NewRequest(http.MethodGet, testUrl, strings.NewReader(``))
		req.Header.Add("X-B3-Traceid", "a818dadb-948a-4b58-80d2-d8afed7078bc")
		rec := httptest.NewRecorder()

		echoContext := echoServer.NewContext(req, rec)
		echoContext.SetParamNames("document")
		echoContext.SetParamValues(document)

		err := controller.GetCustomer(echoContext)

		responseBody := rec.Result().Body
		responseStatusCode := rec.Result().StatusCode

		responseBodyAsString, _ := io.ReadAll(responseBody)
		expectedResponse := "{\"document\":\"98932673667\",\"name\":\"Name\",\"phone\":\"11999999999\",\"email\":\"\"}\n"

		assert.Nil(t, err)
		assert.Equal(t, responseStatusCode, http.StatusOK)
		assert.Equal(t, string(responseBodyAsString), expectedResponse)

		getCustomerPortMock.AssertExpectations(t)
		getCustomerPortMock.AssertCalled(t, "GetCustomer", mock.Anything, document)

	})

	t.Run(`should get no content`, func(t *testing.T) {

		createCustomerPortMock := mocks.NewCreateCustomerPort(t)
		getCustomerPortMock := mocks.NewGetCustomerPort(t)
		controller := Controller{
			CreateCustomerUseCase: createCustomerPortMock,
			GetCustomerUseCase:    getCustomerPortMock,
		}

		document := "98932673667"

		echoServer := echo.New()
		testUrl := fmt.Sprintf("http://localhost:8080/v1/customers/%s", document)

		getCustomerPortMock.On("GetCustomer", mock.Anything, document).Return(nil, nil)

		req, _ := http.NewRequest(http.MethodGet, testUrl, strings.NewReader(``))
		req.Header.Add("X-B3-Traceid", "a818dadb-948a-4b58-80d2-d8afed7078bc")
		rec := httptest.NewRecorder()

		echoContext := echoServer.NewContext(req, rec)
		echoContext.SetParamNames("document")
		echoContext.SetParamValues(document)

		err := controller.GetCustomer(echoContext)

		responseBody := rec.Result().Body
		responseStatusCode := rec.Result().StatusCode

		responseBodyAsString, _ := io.ReadAll(responseBody)
		expectedResponse := "null\n"

		assert.Nil(t, err)
		assert.Equal(t, responseStatusCode, http.StatusNoContent)
		assert.Equal(t, string(responseBodyAsString), expectedResponse)

		getCustomerPortMock.AssertExpectations(t)
		getCustomerPortMock.AssertCalled(t, "GetCustomer", mock.Anything, document)

	})

	t.Run(`should create customer`, func(t *testing.T) {

		createCustomerPortMock := mocks.NewCreateCustomerPort(t)
		getCustomerPortMock := mocks.NewGetCustomerPort(t)
		controller := Controller{
			CreateCustomerUseCase: createCustomerPortMock,
			GetCustomerUseCase:    getCustomerPortMock,
		}

		echoServer := echo.New()
		testUrl := fmt.Sprintf("http://localhost:8080/v1/customers")

		createCustomerPortMock.On("AddCustomer", mock.Anything, mock.Anything).Return(nil)

		req, _ := http.NewRequest(http.MethodPost, testUrl, strings.NewReader(`
			{
				"name": "Name",
				"phone": "11999999999",
				"document": "05251270976"
			}
		`))
		req.Header.Add("X-B3-Traceid", "a818dadb-948a-4b58-80d2-d8afed7078bc")
		rec := httptest.NewRecorder()

		echoContext := echoServer.NewContext(req, rec)
		err := controller.AddCustomer(echoContext)

		responseBody := rec.Result().Body
		responseStatusCode := rec.Result().StatusCode

		responseBodyAsString, _ := io.ReadAll(responseBody)
		expectedResponse := "null\n"

		assert.Nil(t, err)
		assert.Equal(t, responseStatusCode, http.StatusOK)
		assert.Equal(t, string(responseBodyAsString), expectedResponse)

		createCustomerPortMock.AssertExpectations(t)
		createCustomerPortMock.AssertCalled(t, "AddCustomer", mock.Anything, mock.Anything)

	})

	t.Run(`should return error when document is invalid`, func(t *testing.T) {

		createCustomerPortMock := mocks.NewCreateCustomerPort(t)
		getCustomerPortMock := mocks.NewGetCustomerPort(t)
		controller := Controller{
			CreateCustomerUseCase: createCustomerPortMock,
			GetCustomerUseCase:    getCustomerPortMock,
		}

		echoServer := echo.New()
		testUrl := fmt.Sprintf("http://localhost:8080/v1/customers")

		req, _ := http.NewRequest(http.MethodPost, testUrl, strings.NewReader(`
			{
				"name": "Name",
				"phone": "11999999999",
				"document": "12345678910"
			}
		`))
		req.Header.Add("X-B3-Traceid", "a818dadb-948a-4b58-80d2-d8afed7078bc")
		rec := httptest.NewRecorder()

		echoContext := echoServer.NewContext(req, rec)
		err := controller.AddCustomer(echoContext)

		responseBody := rec.Result().Body
		responseStatusCode := rec.Result().StatusCode

		responseBodyAsString, _ := io.ReadAll(responseBody)
		expectedResponse := "{\"code\":400,\"message\":\"INVALID_DATA\"}\n"

		assert.Nil(t, err)
		assert.Equal(t, responseStatusCode, http.StatusBadRequest)
		assert.Equal(t, string(responseBodyAsString), expectedResponse)

		createCustomerPortMock.AssertNotCalled(t, "AddCustomer", mock.Anything, mock.Anything)

	})

	t.Run(`should return error when trying to save customer`, func(t *testing.T) {

		createCustomerPortMock := mocks.NewCreateCustomerPort(t)
		getCustomerPortMock := mocks.NewGetCustomerPort(t)
		controller := Controller{
			CreateCustomerUseCase: createCustomerPortMock,
			GetCustomerUseCase:    getCustomerPortMock,
		}

		echoServer := echo.New()
		testUrl := fmt.Sprintf("http://localhost:8080/v1/customers")

		createCustomerPortMock.On("AddCustomer", mock.Anything, mock.Anything).Return(errors.New("some error"))

		req, _ := http.NewRequest(http.MethodPost, testUrl, strings.NewReader(`
			{
				"name": "Name",
				"phone": "11999999999",
				"document": "05251270976"
			}
		`))
		req.Header.Add("X-B3-Traceid", "a818dadb-948a-4b58-80d2-d8afed7078bc")
		rec := httptest.NewRecorder()

		echoContext := echoServer.NewContext(req, rec)
		err := controller.AddCustomer(echoContext)

		responseBody := rec.Result().Body
		responseStatusCode := rec.Result().StatusCode

		responseBodyAsString, _ := io.ReadAll(responseBody)
		expectedResponse := "{\"code\":500,\"message\":\"some error\"}\n"

		assert.Nil(t, err)
		assert.Equal(t, responseStatusCode, http.StatusInternalServerError)
		assert.Equal(t, string(responseBodyAsString), expectedResponse)

		createCustomerPortMock.AssertExpectations(t)
		createCustomerPortMock.AssertCalled(t, "AddCustomer", mock.Anything, mock.Anything)

	})

}
