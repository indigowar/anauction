package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"

	"github.com/indigowar/anauction/handlers/index"
	"github.com/indigowar/anauction/handlers/login"
	"github.com/indigowar/anauction/handlers/signin"
)

type SetupSettings struct {
	Logger *slog.Logger
}

func Setup(router *echo.Echo, settings SetupSettings) {
	router.Use(slogecho.New(settings.Logger))
	router.Use(middleware.Recover())

	router.Static("/static", "/assets/")

	router.GET("/", index.Page())

	{
		group := router.Group("/auth")

		group.GET("/login", login.Page())
		group.POST("/login", login.HandleRequest())

		group.GET("/signin", signin.Page())
		group.POST("/signin", signin.HandleRequest())
	}
}
