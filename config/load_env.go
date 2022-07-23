package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func LoadENV() {
	// Set config defaults
	viper.SetDefault("DEBUG", true)
	viper.SetDefault("SERVER_ADDR", "127.0.0.1:3000")
	viper.SetDefault("CACHE_DURATION", 2000)

	// Load config file
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // Look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal().Err(err).Msg("fatal error config file")
	}
}
