package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

func main() {
	app := vodka.DefaultRouter()

	app.GET("/ping", func(c *vodka.Context) {
		c.JSON(http.StatusOK, vodka.M{"message": "pong"})
	})

	// Simulates a slow endpoint like a DB query or file upload.
	// When Ctrl+C or SIGTERM arrives, this request will be
	// allowed to finish before the server closes.
	app.GET("/slow", func(c *vodka.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, vodka.M{"message": "finished"})
	})

	// Pass RunConfig with GracefulShutdown: true to enable graceful shutdown.
	// Without RunConfig, Run behaves exactly as before.
	if err := app.Run(":8080", vodka.RunConfig{
		GracefulShutdown: true,
		GracefulTimeout:  15 * time.Second,
	}); err != nil {
		log.Fatal("shutdown error:", err)
	}

	log.Println("Server exited cleanly.")
}
