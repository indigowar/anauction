package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/templates"
)

func Index(c echo.Context) error {
	return templates.Index().Render(c.Request().Context(), c.Response().Writer)
}
