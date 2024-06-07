package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/handlers/index"
)

type SetupSettings struct{}

func Setup(router *echo.Echo, settings SetupSettings) {
	router.Static("/static", "/assets/")

	router.GET("/", index.Page())
}
