package handlers

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"github.com/spazzymoto/echo-scs-session"

	"github.com/indigowar/anauction/domain/service"
	"github.com/indigowar/anauction/handlers/index"
	"github.com/indigowar/anauction/handlers/login"
	"github.com/indigowar/anauction/handlers/logout"
	"github.com/indigowar/anauction/handlers/signin"
)

type SetupSettings struct {
	Logger         *slog.Logger
	SessionManager *scs.SessionManager
	Auth           *service.Auth
}

func Setup(router *echo.Echo, settings SetupSettings) {
	router.Use(session.LoadAndSave(settings.SessionManager))
	router.Use(slogecho.New(settings.Logger))
	router.Use(middleware.Recover())

	router.Static("/static", "/assets/")

	router.GET("/", index.Page(settings.SessionManager))

	{
		group := router.Group("/auth")

		group.GET("/login", login.Page(settings.SessionManager))
		group.POST("/login", login.HandleRequest(settings.SessionManager, settings.Auth))

		group.GET("/signin", signin.Page(settings.SessionManager))
		group.POST("/signin", signin.HandleRequest(settings.Auth, settings.SessionManager))

		group.POST("/logout", logout.HandleRequest(settings.SessionManager))
	}
}
