package swagger

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ISwagger interface {
	RegisterEchoRoutes(e *echo.Echo)
}
type Swagger struct{}

func (c *Swagger) RegisterEchoRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
