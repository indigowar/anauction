package login

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/domain/service"
)

func Page(sm *scs.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := sm.GetBytes(c.Request().Context(), "user-id"); id != nil {
			// The user already logged in, so just a redirect
			c.Response().Header().Add("HX-Redirect", "/")
			return c.NoContent(http.StatusBadRequest)
		}

		return login().Render(c.Request().Context(), c.Response().Writer)
	}
}

func HandleRequest(sm *scs.SessionManager, auth *service.Auth) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := sm.GetBytes(c.Request().Context(), "user-id"); id != nil {
			// The user already logged in, so just a redirect
			c.Response().Header().Add("HX-Redirect", "/")
			return c.NoContent(http.StatusBadRequest)
		}

		emailStr := c.FormValue("email")
		password := c.FormValue("password")

		email, err := mail.ParseAddress(emailStr)
		if err != nil {
			return formWithAnError(
				emailStr, password,
				"Email is invalid",
			).Render(c.Request().Context(), c.Response().Writer)
		}

		id, err := auth.Login(c.Request().Context(), email, password)
		if err != nil {
			errorMsg := "Internal Server error occurred, try again later."

			if errors.Is(err, service.ErrInvalidCredentials) {
				errorMsg = "Provided credentials are invalid"
			}

			return formWithAnError(
				emailStr, password, errorMsg,
			).Render(
				c.Request().Context(),
				c.Response().Writer,
			)
		}

		sm.Put(c.Request().Context(), "user-id", id)

		// Through HTMX make a redirect to home
		c.Response().Header().Add("HX-Redirect", "/")
		return c.NoContent(http.StatusNotImplemented)
	}
}
