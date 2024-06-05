package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/templates"
)

// ServeLoginPage - serves login form
func ServeLoginPage(handle string, signIn string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return templates.Login(handle, signIn).Render(c.Request().Context(), c.Response().Writer)
	}
}
