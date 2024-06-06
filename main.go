package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// TODO: Init storage
	// TODO: Init services

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	router := echo.New()

	handlers.SetupRouter(router, handlers.SetupArgs{
		Logger:         logger,
		SessionManager: sessionManager,
	})

	// TODO: add other handlers

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("stopping the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to stop gracefully,%s\nShutting down by force.", err)
	}

	log.Println("Server is stopped gracefully")
}
