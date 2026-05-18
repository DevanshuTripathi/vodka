package main

import (
	"log"
	"time"

	"github.com/DevanshuTripathi/vodka"
)

func main() {
	app := vodka.DefaultRouter()

	app.GET("/ping", func(c *vodka.Context) {
		c.JSON(200, vodka.M{"message": "pong"})
	})

	// Server starts normally.
	// Press Ctrl+C (SIGINT) or send SIGTERM and it will:
	//   1. Stop accepting new connections
	//   2. Wait up to 10 seconds for active requests to finish
	//   3. Exit cleanly
	if err := app.RunGraceful(":8080", 10*time.Second); err != nil {
		log.Fatal(err)
	}
}
