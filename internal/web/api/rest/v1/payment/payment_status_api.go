package payment

import (
	"bytes"
	"encoding/json"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/web/api/middleware"
	"hamburgueria/internal/web/api/rest/v1/payment/request"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Webhook struct {
	CreatePaymentStatusUseCase input.CreatePaymentStatusPort
}

func (wh *Webhook) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/webhook/payments_status",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", wh.AddPaymentStatus)

}

// AddPaymentStatus
// @Summary     Add payment_status
// @Description Add payment_status
// @Produce      json
// @Param 		 request 	   body   request.CreatePaymentStatusRequest true "Request Body"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.OrderResponse
// @Router       /v1/payments_status [post]
func (wh *Webhook) AddPaymentStatus(ctx echo.Context) error {

	payloadBuffer := new(bytes.Buffer)
	_, err := payloadBuffer.ReadFrom(ctx.Request().Body)
	if err != nil {
		return err
	}
	err = ctx.Request().Body.Close()
	if err != nil {
		return err
	}

	var createPaymentStatusRequest request.CreatePaymentStatusRequest
	errJson := json.Unmarshal(payloadBuffer.Bytes(), &createPaymentStatusRequest)

	if errJson != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": "UNMARSHAL_ERROR",
		})
	}

	errCreated := wh.CreatePaymentStatusUseCase.AddPaymentStatus(ctx.Request().Context(), createPaymentStatusRequest.ToCommand())
	if errCreated != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, nil)
}
