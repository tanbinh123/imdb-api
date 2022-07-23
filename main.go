package main

import (
	"github.com/Scrip7/imdb-api/config"
	"github.com/Scrip7/imdb-api/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// @title        IMDb-API
// @description  Cross-platform microservice to scrape the IMDb website.
// @accept       json
// @produce      json
// @contact.url  https://github.com/Scrip7/imdb-api
// @contact.name Repository
// @license.name MIT license
// @license.url  https://github.com/Scrip7/imdb-api/blob/main/LICENSE
func main() {
	config.LoadENV()
	config.RegisterLogger()

	app := server.GetFiberApp()

	log.Info().Str("addr", viper.GetString("SERVER_ADDR")).Msg("Starting the HTTP server")
	log.Fatal().Err(app.Listen(viper.GetString("SERVER_ADDR"))).Msg("Failed to start the HTTP server")
}
