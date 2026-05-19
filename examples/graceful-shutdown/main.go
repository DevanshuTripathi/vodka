package main

import (
	"log"

	"github.com/DevanshuTripathi/vodka"
)

func main() {
	app := vodka.DefaultRouter()

	app.GET("/ping", func(c *vodka.Context) {
		c.JSON(200, vodka.M{"message": "pong"})
	})

	// Start the server.
	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
