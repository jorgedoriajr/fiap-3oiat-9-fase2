package order

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/application/api/middleware"
	"hamburgueria/internal/modules/order/domain/request"
	"hamburgueria/internal/modules/order/domain/response"
	"hamburgueria/internal/modules/order/port/input"
	"net/http"
)

type Controller struct {
	CreateOrderUseCase input.CreateOrderPort
}

func (c *Controller) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/orders",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddOrder)
}

// AddOrder
// @Summary     Add order
// @Description Add order
// @Produce      json
// @Param 		 request 	   body   request.CreateOrder true "Request Body"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      200
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

	result, err := c.CreateOrderUseCase.AddOrder(ctx.Request().Context(), createOrderRequest.ToCommand())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.OrderResponse{
		Amount:      result.Amount,
		PaymentData: result.PaymentData,
	})
}
