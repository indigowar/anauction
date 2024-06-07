package index

import (
	"github.com/labstack/echo/v4"
)

func Page() echo.HandlerFunc {
	return func(c echo.Context) error {
		return index().Render(c.Request().Context(), c.Response().Writer)
	}
}
