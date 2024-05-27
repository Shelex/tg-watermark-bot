package env

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	TgToken     string
	Locale      string
}

var Settings *Config

func ReadEnv() *Config {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BOT")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to read config file: %s", err)
	}

	viper.SetDefault("ENVIRONMENT", "dev")
	viper.SetDefault("LOCALE", "uk")

	Settings = &Config{
		Environment: viper.GetString("ENVIRONMENT"),
		TgToken:     viper.GetString("TELEGRAM_TOKEN"),
		Locale:      viper.GetString("LOCALE"),
	}

	return Settings
}
