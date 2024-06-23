package logout

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
)

func HandleRequest(sm *scs.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := sm.GetString(c.Request().Context(), "user-id"); id != "" {
			// The user already logged in, so just a redirect
			c.Response().Header().Add("HX-Redirect", "/")
			return c.NoContent(http.StatusBadRequest)
		}

		sm.Remove(c.Request().Context(), "user-id")

		return c.Redirect(http.StatusPermanentRedirect, "/")
	}
}
