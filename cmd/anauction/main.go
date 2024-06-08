package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := echo.New()

	handlers.Setup(router, handlers.SetupSettings{
		Logger: logger,
	})

	run(router, logger)
}

func run(router *echo.Echo, logger *slog.Logger) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(
				"server ListenAndServe failed",
				"err", err,
			)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error(
			"Failed to stop the server gracefully",
			"Error", err,
		)
		os.Exit(1)
	}

	logger.Info("Server is stopped")
}
