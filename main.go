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
	const addr = "127.0.0.1:8000"

	r := echo.New()

	r.Static("/static", "./assets/")

	r.GET("/", handlers.Index)
	r.GET("/auth/login", handlers.ServeLoginPage(addr+"/auth/register", addr+"/auth/signin"))

	server := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil || err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shut down: %s", err.Error())
	}

	log.Println("Server is stopped")
}
