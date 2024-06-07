package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		return login().Render(c.Request().Context(), c.Response().Writer)
	}
}

func HandleRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotImplemented)
	}
}
