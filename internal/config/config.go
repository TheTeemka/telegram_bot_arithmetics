package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	LogLevel         string
	TelegramBotToken string
	NumWorkers       int
}

func LoadConfig() *Config {
	cfg := &Config{
		LogLevel:         os.Getenv("LOG_LEVEL"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		NumWorkers:       1,
	}

	if cfg.TelegramBotToken == "" {
		panic("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	strWorkers := os.Getenv("NUM_WORKERS")
	if n, err := strconv.Atoi(strWorkers); err == nil && n > 0 {
		cfg.NumWorkers = n
	}

	return cfg
}
