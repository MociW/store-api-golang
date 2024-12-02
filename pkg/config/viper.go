package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	// Set the configuration file name and type
	config.SetConfigName(".env") // Include the dot if the file name is ".env"
	config.SetConfigType("env")  // Specify file type as "env"

	// Add paths to look for the config file
	config.AddConfigPath("./../") // Parent directory
	config.AddConfigPath("./")    // Current directory

	// Read in the configuration file
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
