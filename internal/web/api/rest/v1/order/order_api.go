package order

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/usecase/result"
	"hamburgueria/internal/web/api/middleware"
	"hamburgueria/internal/web/api/rest/v1/order/presenter"
	"hamburgueria/internal/web/api/rest/v1/order/request"
	"net/http"
	"strconv"
)

type Api struct {
	CreateOrderUseCase input.CreateOrderPort
	ListOrderUseCase   input.ListOrderPort
}

func (c *Api) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/orders",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddOrder)
	group.Add(http.MethodGet, "", c.GetOrders)
	group.Add(http.MethodGet, "/:number", c.GetOrderByNumber)
}

// AddOrder
// @Summary     Add order
// @Description Add order
// @Produce      json
// @Param 		 request 	   body   request.CreateOrder true "Request Body"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.OrderResponse
// @Router       /v1/orders [post]
func (c *Api) AddOrder(ctx echo.Context) error {

	payloadBuffer := new(bytes.Buffer)
	_, err := payloadBuffer.ReadFrom(ctx.Request().Body)
	if err != nil {
		return err
	}
	err = ctx.Request().Body.Close()
	if err != nil {
		return err
	}

	var createOrderRequest request.CreateOrder
	errJson := json.Unmarshal(payloadBuffer.Bytes(), &createOrderRequest)

	if errJson != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": "UNMARSHAL_ERROR",
		})
	}

	orderCreated, err := c.CreateOrderUseCase.AddOrder(ctx.Request().Context(), createOrderRequest.ToCommand())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, presenter.OrderResponseFromResult(*orderCreated))
}

// GetOrders
// @Summary      Get Orders
// @Description  Get Orders
// @Produce      json
// @Param 		 status query string false "Filter Orders by status"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.ListOrderResponse
// @Router       /v1/orders/ [get]
func (c *Api) GetOrders(ctx echo.Context) error {
	//TODO need pagination
	status := ctx.QueryParam("status")

	var resultOrders []result.ListOrderResult
	var err error

	if status == "" {
		resultOrders, err = c.ListOrderUseCase.FindAllOrders(ctx.Request().Context())
	} else {
		resultOrders, err = c.ListOrderUseCase.FindByStatus(ctx.Request().Context(), status)
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if resultOrders == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.ListOrderResponseFromResult(resultOrders))
}

// GetOrderByNumber
// @Summary      Get Order by number
// @Description  Get Order by number
// @Produce      json
// @Param        number    path      string  true  "number"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.ListOrderResponse
// @Router       /v1/products/{number} [get]
func (c *Api) GetOrderByNumber(ctx echo.Context) error {
	numberPathParam := ctx.Param("number")
	if numberPathParam == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "number cannot be empty",
		})
	}
	number, err := strconv.Atoi(numberPathParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "number must be a numeric value",
		})
	}

	order, err := c.ListOrderUseCase.FindByNumber(ctx.Request().Context(), number)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
	}

	if order == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.GetOrderResponseFromResult(*order))
}
