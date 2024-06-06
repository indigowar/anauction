package handlers

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	session "github.com/spazzymoto/echo-scs-session"
)

type SetupArgs struct {
	Logger         *slog.Logger
	SessionManager *scs.SessionManager

	StaticFilesPrefix string
	StaticFilesDir    string
}

func SetupRouter(router *echo.Echo, args SetupArgs) {
	router.Use(session.LoadAndSave(args.SessionManager))

	router.Static(args.StaticFilesPrefix, args.StaticFilesDir)

	router.GET("/", indexPage())

	{
		auth := router.Group("/auth")

		auth.GET("/login", loginPage("/auth/login", "/auth/signin"))
	}
}
