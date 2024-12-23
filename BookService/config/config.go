package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	DefaultLanguage string `mapstructure:"DEFAULT_LANGUAGE"`

	AppName        string `mapstructure:"APP_NAME"`
	AppPort        uint32 `mapstructure:"APP_PORT"`
	AppPrivatePort uint32 `mapstructure:"APP_PRIVATE_PORT"`
	AppHost        string `mapstructure:"APP_HOST"`
	AppBaseUrl     string `mapstructure:"APP_BASEURL"`
	AppBasicAuth   string `mapstructure:"APP_BASIC_AUTH"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	JWTSecret string `mapstructure:"JWT_SECRET"`
}

// LoadConfig loads the configuration from the .env file.
func LoadConfig() (*Config, error) {
	// Set the path to look for the .env file
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Automatically override environment variables
	viper.AutomaticEnv()

	// Replace `.` in environment variable keys with `_` to support nested structures
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, using environment variables: %v", err)
	}

	// Unmarshal into the Config struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
