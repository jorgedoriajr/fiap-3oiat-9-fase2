package ingredient

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"hamburgueria/internal/server/api/middleware"
	"hamburgueria/internal/server/api/rest/v1/ingredient/request"
	"hamburgueria/internal/server/api/rest/v1/ingredient/response"

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
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.IngredientResponse
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

	resultIngredient, err := c.CreateIngredientUseCase.AddIngredient(e.Request().Context(), req.ToCommand())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, response.FromCreateIngredientResult(resultIngredient))
}

// GetIngredientByID
// @Summary     Get Ingredient by id
// @Description Get Ingredient by id
// @Produce      json
// @Param        id    path      string  true  "id"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.IngredientResponse
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
	ingredientResult, err := c.IngredientFinderService.FindIngredientByID(ctx.Request().Context(), ingredientID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if ingredientResult == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, response.FromFindIngredientResult(*ingredientResult))
}

// GetIngredients
// @Summary     Get Ingredients
// @Description Get Ingredients
// @Produce      json
// @Param 		type query string false "Filter Ingredients by type"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} []response.IngredientResponse
// @Router       /v1/ingredients [get]
func (c *Controller) GetIngredients(ctx echo.Context) error {
	ingredientType := ctx.QueryParam("type")

	var resultIngredients []result.FindIngredientResult
	var err error

	if ingredientType != "" {
		resultIngredients, err = c.IngredientFinderService.FindIngredientByType(ctx.Request().Context(), ingredientType)
	} else {
		resultIngredients, err = c.IngredientFinderService.FindAllIngredients(ctx.Request().Context())
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": err.Error(),
		})
	}

	if resultIngredients == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, response.FromFindIngredientsResult(resultIngredients))

}
