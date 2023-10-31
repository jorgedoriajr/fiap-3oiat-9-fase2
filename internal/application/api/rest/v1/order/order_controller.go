package order

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/application/api/middleware"
	"hamburgueria/internal/application/api/rest/v1/order/request"
	"hamburgueria/internal/application/api/rest/v1/order/response"
	"hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/usecase/result"
	"net/http"
)

type Controller struct {
	CreateOrderUseCase input.CreateOrderPort
	ListOrderUseCase   input.ListOrderPort
}

func (c *Controller) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/orders",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddOrder)
	group.Add(http.MethodGet, "", c.GetOrders)
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
func (c *Controller) AddOrder(ctx echo.Context) error {

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
	return ctx.JSON(http.StatusOK, response.OrderResponse{
		Amount:      orderCreated.Amount,
		PaymentData: orderCreated.PaymentData,
	})
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
func (c *Controller) GetOrders(ctx echo.Context) error {

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

	var ordersResponse []response.ListOrderResponse
	for _, order := range resultOrders {
		ordersResponse = append(ordersResponse, response.ListOrderResponse{
			OrderId:    order.OrderId,
			Status:     order.Status,
			Amount:     order.Amount,
			CustomerId: order.CustomerId,
			CreatedAt:  order.CreatedAt,
		})
	}
	return ctx.JSON(http.StatusOK, ordersResponse)
}
