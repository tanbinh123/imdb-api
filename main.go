package main

import (
	"os"
	"time"

	"github.com/spf13/viper"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/Scrip7/imdb-api/routes/title"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set config defaults
	viper.SetDefault("SERVER_ADDR", "127.0.0.1:3000")

	// Load config file
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // Look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal().Msgf("fatal error config file: %v", err)
	}

	// Default level for this example is info, unless debug flag is present
	if viper.GetBool("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Info().Msg("Config loaded")

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Err(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":      false,
				"message": err.Error(),
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
		Expiration: 5 * time.Second,
		// CacheControl: true,
	}))

	t := app.Group("/title/:id")
	t.Get("/", title.Index)
	t.Get("/debug", title.IndexDebug)

	log.Info().Str("addr", viper.GetString("SERVER_ADDR")).Msg("Starting the HTTP server")
	log.Fatal().Err(app.Listen(viper.GetString("SERVER_ADDR"))).Msg("Failed to start the HTTP server")
}
