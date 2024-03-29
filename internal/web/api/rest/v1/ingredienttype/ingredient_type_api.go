package ingredienttype

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/web/api/middleware"
	"hamburgueria/internal/web/api/rest/v1/ingredienttype/presenter"
	"net/http"
)

type Api struct {
	FindIngredientTypeUseCase input.FindIngredientTypeUseCasePort
}

func (c *Api) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/ingredient-types",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodGet, "", c.GetIngredientTypes)
}

// GetIngredientTypes
// @Summary     Get Ingredient types
// @Description Get Ingredient types
// @Produce      json
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.IngredientTypeResponse
// @Router       /v1/ingredient-types [get]
func (c *Api) GetIngredientTypes(ctx echo.Context) error {

	result, err := c.FindIngredientTypeUseCase.FindAll(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if result == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.IngredientTypeResponseFromResult(result))

}
