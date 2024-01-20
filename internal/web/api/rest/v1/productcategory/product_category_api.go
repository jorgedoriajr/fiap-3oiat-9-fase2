package productcategory

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/product/ports/input"
	"hamburgueria/internal/web/api/middleware"
	"hamburgueria/internal/web/api/rest/v1/productcategory/presenter"
	"net/http"
)

type Api struct {
	GetProductCategoryUseCase input.GetProductCategoryUseCasePort
}

func (c *Api) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/product-categories",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodGet, "", c.GetProductCategories)
}

// GetProductCategories
// @Summary     Get Product Categories
// @Description Get Product Categories
// @Produce      json
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.ProductCategoryResponse
// @Router       /v1/product-category [get]
func (c *Api) GetProductCategories(ctx echo.Context) error {

	result, err := c.GetProductCategoryUseCase.FindAll(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if result == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.ProductCategoriesResponseFromResult(result))

}
