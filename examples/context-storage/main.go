package main

import (
	"log"

	"github.com/DevanshuTripathi/vodka"
)

func addRequestUser() vodka.HandlerFunc {
	return func(c *vodka.Context) {
		user := c.Query("user")
		if user == "" {
			user = "guest"
		}

		c.Set("user", user)
		c.Next()
	}
}

func main() {
	app := vodka.DefaultRouter()
	app.Use(addRequestUser())

	app.GET("/profile", func(c *vodka.Context) {
		user, _ := c.Get("user")
		c.JSON(200, vodka.M{
			"message": "context value loaded",
			"user":    user,
		})
	})

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
