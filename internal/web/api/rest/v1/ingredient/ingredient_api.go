package ingredient

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"hamburgueria/internal/web/api/middleware"
	"hamburgueria/internal/web/api/rest/v1/ingredient/presenter"
	"hamburgueria/internal/web/api/rest/v1/ingredient/request"
	"strconv"

	"net/http"
)

type Api struct {
	CreateIngredientUseCase input.CreateIngredientUseCasePort
	FindIngredientUseCase   input.FindIngredientUseCasePort
}

func (c *Api) RegisterEchoRoutes(e *echo.Echo) {
	group := e.Group("/v1/ingredients",
		middleware.GetTraceCallsMiddlewareFunc(),
		middleware.GetLogCallsMiddlewareFunc(),
	)
	group.Add(http.MethodPost, "", c.AddIngredient)
	group.Add(http.MethodGet, "", c.GetIngredients)
	group.Add(http.MethodGet, "/:number", c.GetIngredientByNumber)
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
func (c *Api) AddIngredient(e echo.Context) error {
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
	return e.JSON(http.StatusOK, presenter.CreateIngredientResponseFromResult(*resultIngredient))
}

// GetIngredientByNumber
// @Summary     Get Ingredient by number
// @Description Get Ingredient by number
// @Produce      json
// @Param        number    path      string  true  "number"
// @Failure      400 {object} v1.ErrorResponse
// @Failure      401 {object} v1.ErrorResponse
// @Failure      404 {object} v1.ErrorResponse
// @Failure      503 {object} v1.ErrorResponse
// @Success      200 {object} response.IngredientResponse
// @Router       /v1/ingredients/{number} [get]
func (c *Api) GetIngredientByNumber(ctx echo.Context) error {
	numberParam := ctx.Param("number")
	if numberParam == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "number cannot be empty",
		})
	}

	number, err := strconv.Atoi(numberParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"code":    400,
			"message": "number needs to be a numeric value",
		})
	}

	ingredientResult, err := c.FindIngredientUseCase.FindIngredientByNumber(ctx.Request().Context(), number)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"code":    500,
			"message": err.Error(),
		})
	}

	if ingredientResult == nil {
		return ctx.JSON(http.StatusNoContent, nil)
	}

	return ctx.JSON(http.StatusOK, presenter.FindIngredientResponseFromResult(*ingredientResult))
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
func (c *Api) GetIngredients(ctx echo.Context) error {
	ingredientType := ctx.QueryParam("type")

	var resultIngredients []result.FindIngredientResult
	var err error

	if ingredientType != "" {
		resultIngredients, err = c.FindIngredientUseCase.FindIngredientByType(ctx.Request().Context(), ingredientType)
	} else {
		resultIngredients, err = c.FindIngredientUseCase.FindAllIngredients(ctx.Request().Context())
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

	return ctx.JSON(http.StatusOK, presenter.FindIngredientsResponseFromResult(resultIngredients))

}
