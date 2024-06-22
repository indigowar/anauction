package signin

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/domain/service"
)

func Page(sm *scs.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := sm.GetString(c.Request().Context(), "user-id"); id != "" {
			// The user already logged in, so just a redirect
			c.Response().Header().Add("HX-Redirect", "/")
			return c.NoContent(http.StatusBadRequest)
		}

		return signIn().Render(c.Request().Context(), c.Response().Writer)
	}
}

func HandleRequest(auth *service.Auth, sm *scs.SessionManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		if id := sm.GetString(c.Request().Context(), "user-id"); id != "" {
			// The user already logged in, so just a redirect
			c.Response().Header().Add("HX-Redirect", "/")
			return c.NoContent(http.StatusBadRequest)
		}

		name := c.FormValue("name")
		password := c.FormValue("password")
		emailStr := c.FormValue("email")

		email, err := mail.ParseAddress(emailStr)
		if err != nil {
			return formWithInvalidEmail(
				name, emailStr, password,
				"Email is invalid",
			).Render(c.Request().Context(), c.Response().Writer)
		}

		id, err := auth.SignIn(c.Request().Context(), name, email, password)
		if err != nil {
			var duplicationErr *service.DuplicationError
			if errors.As(err, &duplicationErr) {
				return formWithFormError(
					name, emailStr, password,
					fmt.Sprintf("The %s with provided %s already exists", duplicationErr.Object, duplicationErr.Field),
				).Render(c.Request().Context(), c.Response().Writer)
			}

			return formWithFormError(
				name, emailStr, password,
				"Internal server error occurred, try again later",
			).Render(c.Request().Context(), c.Response().Writer)
		}

		sm.Put(c.Request().Context(), "user-id", id.String())

		c.Response().Header().Add("HX-Redirect", "/")
		return c.NoContent(http.StatusOK)
	}
}
