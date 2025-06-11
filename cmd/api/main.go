package main

import (
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/TheTeemka/telegram_bot_arithmetics/internal/config"
	"github.com/TheTeemka/telegram_bot_arithmetics/internal/telegram"
)

func main() {
	cfg := config.LoadConfig()

	setSlog(cfg.LogLevel)

	bot, err := telegram.NewBot(cfg.TelegramBotToken, cfg.NumWorkers)
	if err != nil {
		slog.Error("Failed to create bot", "error", err)
	}

	bot.Start()
}

func setSlog(logLevel string) {
	var level slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info", "":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		log.Fatal("Invalid log level: " + logLevel)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
}
