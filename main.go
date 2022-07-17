package main

import (
	"log"
	"time"

	"github.com/Scrip7/imdb-api/routes/title"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		},
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Cache middleware
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			// TODO: replace with header
			return c.Query("refresh") == "true"
		},
		// TODO: load from config
		Expiration: 1 * time.Second,
		// CacheControl: true,
	}))

	t := app.Group("/title/:id")
	t.Get("/", title.Index)

	// TODO: load url from config
	log.Fatal(app.Listen("127.0.0.1:3000"))
}
