package ingredienttype

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/server/api/middleware"
	"hamburgueria/internal/server/api/rest/v1/ingredienttype/response"
	"net/http"
)

type Controller struct {
	IngredientTypeFinderService input.IngredientTypeFinderServicePort
}

func (c *Controller) RegisterEchoRoutes(e *echo.Echo) {
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
func (c *Controller) GetIngredientTypes(ctx echo.Context) error {

	result, err := c.IngredientTypeFinderService.FindAllIngredientType(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if result == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	var ingredientTypeResponse []response.IngredientTypeResponse
	for _, ingredientType := range result {
		ingredientTypeResponse = append(ingredientTypeResponse, response.IngredientTypeResponse{Name: ingredientType.Name})
	}
	return ctx.JSON(http.StatusOK, ingredientTypeResponse)

}
