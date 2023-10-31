package product

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/application/api/middleware"
	"hamburgueria/internal/application/api/rest/v1/product/request"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/ports/input"
	"net/http"
)

type Controller struct {
	CreateProductUseCase input.CreateProductUseCasePort
	ProductFinderService input.ProductFinderServicePort
}

func (c *Controller) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/products",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddProduct)
	group.Add(http.MethodGet, "", c.GetProducts)
	group.Add(http.MethodGet, "/:productId", c.GetProductById)
}

// AddProduct
// @Summary     Add Product
// @Description Add Product
// @Produce      json
// @Param 		 request 	   body   request.CreateCustomerCommand true "Request Body"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      201
// @Router       /v1/products [post]
func (c *Controller) AddProduct(e echo.Context) error {
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

// GetProductById
// @Summary     Get Product by id
// @Description Get Product by id
// @Produce      json
// @Param        id    path      string  true  "id"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      200
// @Router       /v1/products/{productID} [get]
func (c *Controller) GetProductById(ctx echo.Context) error {
	id := ctx.Param("productId")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "id cannot be empty",
		})
	}
	productID := uuid.MustParse(id)
	response, err := c.ProductFinderService.FindByID(ctx.Request().Context(), productID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if response == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetProducts
// @Summary     Get Products
// @Description Get Products
// @Produce      json
// @Param 		category query string false "Filter products by category"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      200 {object} []entity.ProductEntity
// @Router       /v1/products [get]
func (c *Controller) GetProducts(ctx echo.Context) error {
	category := ctx.QueryParam("category")

	var response []entity.ProductEntity
	var err error

	if category != "" {
		response, err = c.ProductFinderService.FindByCategory(ctx.Request().Context(), category)
	} else {
		response, err = c.ProductFinderService.FindAllProducts(ctx.Request().Context())
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if response == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, response)

}
