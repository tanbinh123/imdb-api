package main

import (
	"time"

	"github.com/Scrip7/imdb-api/services/imdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	app := fiber.New()

	// Cache middleware
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		// TODO: load from config
		Expiration: 1 * time.Second,
		// CacheControl: true,
	}))

	// TODO: use routes package
	app.Get("/title/:id", func(c *fiber.Ctx) error {
		// TODO: validate params
		res, err := imdb.Title(c.Params("id"))
		if err != nil {
			// TODO: global error handling middleware
			return err
		}
		return c.JSONP(res)
	})

	// TODO: load url from config
	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}
