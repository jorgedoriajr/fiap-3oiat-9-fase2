package customer

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/customer/port/input"
	"hamburgueria/internal/server/api/middleware"
	"hamburgueria/internal/server/api/rest/presenter"
	"hamburgueria/internal/server/api/rest/v1/customer/request"
	"hamburgueria/pkg/validation"
	"net/http"
)

type Controller struct {
	CreateCustomerUseCase input.CreateCustomerPort
	GetCustomerUseCase    input.GetCustomerPort
}

func (c *Controller) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/customers",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodGet, "/:document", c.GetCustomer)
	group.Add(http.MethodPost, "", c.AddCustomer)
}

// GetCustomer
// @Summary     Get Customer by document
// @Description Get Customer by document
// @Produce      json
// @Param        document    path      string  true  "Document"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.Customer
// @Router       /v1/customers/{document} [get]
func (c *Controller) GetCustomer(ctx echo.Context) error {
	document := ctx.Param("document")

	response, err := c.GetCustomerUseCase.GetCustomer(ctx.Request().Context(), document)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if response == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.ToGetCustomerResponse(*response))
}

// AddCustomer
// @Summary     Add customer
// @Description Add customer
// @Produce      json
// @Param 		 request 	   body   request.CreateCustomer true "Request Body"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200
// @Router       /v1/customers [post]
func (c *Controller) AddCustomer(ctx echo.Context) error {

	payloadBuffer := new(bytes.Buffer)
	_, err := payloadBuffer.ReadFrom(ctx.Request().Body)
	if err != nil {
		return err
	}
	err = ctx.Request().Body.Close()
	if err != nil {
		return err
	}

	var createCustomerRequest request.CreateCustomer
	errJson := json.Unmarshal(payloadBuffer.Bytes(), &createCustomerRequest)

	if errJson != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": "UNMARSHAL_ERROR",
		})
	}

	isValid := validation.ValidateCPF(createCustomerRequest.Document)
	if !isValid {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": "INVALID_DATA",
		})
	}

	err = c.CreateCustomerUseCase.AddCustomer(ctx.Request().Context(), createCustomerRequest.ToCommand())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, nil)
}
