package product

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/application/api/middleware"
	"hamburgueria/internal/application/api/rest/v1/product/request"
	"hamburgueria/internal/modules/product/ports/input"
	"net/http"
)

type CategoryController struct {
	ProductFinderService input.ProductFinderServicePort
}

func (c *CategoryController) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/products/categories",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddProduct)
	group.Add(http.MethodGet, "", c.GetProducts)
	group.Add(http.MethodGet, "/:productId", c.GetProductById)
}

// AddProductCategory
// @Summary     Add Product Category
// @Description Add Product Category
// @Produce      json
// @Param 		 request 	   body   request.CreateCustomerCommand true "Request Body"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      201
// @Router       /v1/products/categories [post]
func (c *CategoryController) AddProductCategory(e echo.Context) error {
	req := new(request.CreateProductRequest)

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "UNMARSHAL_ERROR",
		})
	}

	if errBind := e.Bind(req); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "UNMARSHAL_ERROR",
		})
	}

	result, err := c.CreateProductUseCase.AddProduct(e.Request().Context(), req.ToCommand())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, result)
}
