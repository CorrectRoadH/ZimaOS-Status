package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (v *StatusRoute) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello World")
}
