package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/indigowar/anauction/handlers"
)

func main() {
	router := echo.New()

	handlers.Setup(router, handlers.SetupSettings{})

	run(router)
}

func run(router *echo.Echo) {
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to stop server gracefully: %s\n Forcing to shut down", err)
	}

	log.Println("Server is stopped")
}
