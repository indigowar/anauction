package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/handlers/index"
	"github.com/indigowar/anauction/handlers/login"
)

type SetupSettings struct {
	Logger *slog.Logger
}

func Setup(router *echo.Echo, settings SetupSettings) {
	router.Static("/static", "/assets/")

	router.GET("/", index.Page())

	{
		group := router.Group("/auth")

		group.GET("/login", login.Page())
		group.POST("/login", login.HandleRequest())
	}
}