package templates

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, status int, t templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().Writer.WriteHeader(status)
	err := t.Render(c.Request().Context(), c.Response().Writer)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to render response template")
	}
	return nil
}
