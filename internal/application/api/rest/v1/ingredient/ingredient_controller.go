package ingredient

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/application/api/middleware"
	"hamburgueria/internal/application/api/rest/v1/ingredient/request"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/usecase/result"

	"net/http"
)

type Controller struct {
	CreateIngredientUseCase input.CreateIngredientUseCasePort
	IngredientFinderService input.IngredientFinderServicePort
}

func (c *Controller) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/ingredients",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddIngredient)
	group.Add(http.MethodGet, "", c.GetIngredients)
	group.Add(http.MethodGet, "/:ingredientID", c.GetIngredientByID)
}

// AddIngredient
// @Summary     Add Ingredient
// @Description Add Ingredient
// @Produce      json
// @Param 		 request 	   body   request.CreateIngredientRequest true "Request Body"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      201
// @Router       /v1/ingredients [post]
func (c *Controller) AddIngredient(e echo.Context) error {
	req := new(request.CreateIngredientRequest)

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

	result, err := c.CreateIngredientUseCase.AddIngredient(e.Request().Context(), req.ToCommand())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, result.ToResponse())
}

// GetIngredientByID
// @Summary     Get Ingredient by id
// @Description Get Ingredient by id
// @Produce      json
// @Param        id    path      string  true  "id"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      200
// @Router       /v1/ingredients/{ingredientID} [get]
func (c *Controller) GetIngredientByID(ctx echo.Context) error {
	id := ctx.Param("ingredientID")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "id cannot be empty",
		})
	}
	ingredientID := uuid.MustParse(id)
	response, err := c.IngredientFinderService.FindIngredientByID(ctx.Request().Context(), ingredientID)
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

// GetIngredients
// @Summary     Get Ingredients
// @Description Get Ingredients
// @Produce      json
// @Param 		type query string false "Filter Ingredients by type"
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Failure      503 {object} model.ErrorResponse
// @Success      200 {object} []entity.IngredientType
// @Router       /v1/ingredients [get]
func (c *Controller) GetIngredients(ctx echo.Context) error {
	ingredientType := ctx.QueryParam("type")

	var response []result.FindIngredientResult
	var err error

	if ingredientType != "" {
		response, err = c.IngredientFinderService.FindIngredientByType(ctx.Request().Context(), ingredientType)
	} else {
		response, err = c.IngredientFinderService.FindAllIngredients(ctx.Request().Context())
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
