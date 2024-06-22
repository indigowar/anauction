package index

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func Page(sm *scs.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		isLoggedIn := sm.GetString(c.Request().Context(), "user-id") != ""

		return index(isLoggedIn).Render(c.Request().Context(), c.Response().Writer)
	}
}
