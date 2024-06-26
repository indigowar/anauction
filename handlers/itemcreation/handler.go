package itemcreation

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/templates"
)

func Form() echo.HandlerFunc {
	return func(c echo.Context) error {
		return templates.Render(c, http.StatusOK, formPage())
	}
}

func HandleRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("not implemented")
	}
}
