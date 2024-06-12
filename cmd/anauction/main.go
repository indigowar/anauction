package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/domain/service"
	"github.com/indigowar/anauction/handlers"
	"github.com/indigowar/anauction/storage/postgres"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if err := godotenv.Load(); err != nil {
		logger.Error("Failed to load env file", "Err", err)
		os.Exit(1)
	}

	dbConn, err := connectToDB()
	if err != nil {
		logger.Error("Failed to connect to database", "Err", err)
		os.Exit(1)
	}

	userStorage := postgres.NewUserStorage(dbConn)

	authService := service.NewAuth(logger, userStorage)

	sm := scs.New()
	sm.Lifetime = 24 * time.Hour

	router := echo.New()
	handlers.Setup(router, handlers.SetupSettings{
		Logger:         logger,
		SessionManager: sm,
		Auth:           &authService,
	})

	run(router, logger)
}

func connectToDB() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	return pgx.Connect(context.Background(), url)
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
