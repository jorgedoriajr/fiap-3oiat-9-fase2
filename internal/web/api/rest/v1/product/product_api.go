package product

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/product/ports/input"
	"hamburgueria/internal/modules/product/usecase/result"
	"hamburgueria/internal/web/api/middleware"
	"hamburgueria/internal/web/api/rest/v1/product/presenter"
	"hamburgueria/internal/web/api/rest/v1/product/request"
	"net/http"
	"strconv"
)

type Api struct {
	CreateProductUseCase  input.CreateProductUseCasePort
	UpdatedProductUseCase input.UpdateProductUseCasePort
	FindProductUseCase    input.FindProductUseCasePort
	DeleteProductUseCase  input.DeleteProductUseCasePort
}

func (c *Api) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/products",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddProduct)
	group.Add(http.MethodGet, "", c.GetProducts)
	group.Add(http.MethodGet, "/:number", c.GetProductByNumber)
	group.Add(http.MethodDelete, "/:number", c.InactiveProductByNumber)
	group.Add(http.MethodPatch, "/:number", c.UpdateProduct)

}

// AddProduct
// @Summary     Add Product
// @Description Add Product
// @Produce      json
// @Param 		 request 	   body   request.ProductRequest true "Request Body"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.ProductCreatedResponse
// @Router       /v1/products [post]
func (c *Api) AddProduct(e echo.Context) error {
	req := new(request.ProductRequest)

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

	resultProduct, err := c.CreateProductUseCase.AddProduct(e.Request().Context(), req.ToCommand())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, presenter.ProductCreatedResponseFromResult(*resultProduct))
}

// UpdateProduct
// @Summary     Update Product
// @Description Update Product
// @Produce      json
// @Param 		 request 	   body   request.UpdateProductRequest true "Request Body"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200
// @Router       /v1/products/{number} [patch]
func (c *Api) UpdateProduct(ctx echo.Context) error {

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
			"message": "number needs to be a numeric value",
		})
	}

	req := new(request.UpdateProductRequest)
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "UNMARSHAL_ERROR",
		})
	}

	if errBind := ctx.Bind(req); errBind != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "UNMARSHAL_ERROR",
		})
	}

	req.Number = number

	err = c.UpdatedProductUseCase.UpdateProduct(
		ctx.Request().Context(), req.ToCommand(),
	)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, nil)
}

// GetProductByNumber
// @Summary     Get Product by number
// @Description Get Product by number
// @Produce      json
// @Param        number    path      string  true  "number"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.FindProductWithIngredients
// @Router       /v1/products/{number} [get]
func (c *Api) GetProductByNumber(ctx echo.Context) error {
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

	resultProduct, err := c.FindProductUseCase.FindByNumber(ctx.Request().Context(), number)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if resultProduct == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.ProductResponseFromResult(*resultProduct))
}

// InactiveProductByNumber
// @Summary     Delete Product by number
// @Description Delete Product by number
// @Produce      json
// @Param        number    path      string  true  "number"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200
// @Router       /v1/products/{number} [delete]
func (c *Api) InactiveProductByNumber(ctx echo.Context) error {
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
			"message": "number needs to be a numeric value",
		})
	}
	err = c.DeleteProductUseCase.Inactive(ctx.Request().Context(), number)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, nil)
}

// GetProducts
// @Summary     Get Products
// @Description Get Products
// @Produce      json
// @Param 		category query string false "Filter products by category"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.FindProductWithIngredients
// @Router       /v1/products [get]
func (c *Api) GetProducts(ctx echo.Context) error {
	category := ctx.QueryParam("category")

	var resultProducts []result.FindProductResult
	var err error

	if category != "" {
		resultProducts, err = c.FindProductUseCase.FindByCategory(ctx.Request().Context(), category)
	} else {
		resultProducts, err = c.FindProductUseCase.FindAllProducts(ctx.Request().Context())
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if resultProducts == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.ProductsResponseFromResultList(resultProducts))

}
