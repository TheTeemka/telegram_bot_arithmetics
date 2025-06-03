package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	LogLevel         string
	TelegramBotToken string
}

func LoadConfig() *Config {
	cfg := &Config{
		LogLevel:         os.Getenv("LOG_LEVEL"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}

	if cfg.TelegramBotToken == "" {
		panic("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	return cfg
}
