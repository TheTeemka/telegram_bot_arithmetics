package telegram

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/TheTeemka/telegram_bot_arithmetics/internal/arithmetic"
	tapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	BotAPI *tapi.BotAPI
}

func NewBot(token string) (*TelegramBot, error) {
	bot, err := tapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	slog.Info("Telegram Bot started", "bot_name", bot.Self.UserName, "bot_id", bot.Self.ID)
	return &TelegramBot{BotAPI: bot}, nil
}

func (b *TelegramBot) Start() {
	updateConfig := tapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := b.BotAPI.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.HandleCommand(update.Message.Command(), update.Message.Chat.ID)
		} else {
			b.HandleText(update.Message.Text, update.Message.Chat.ID)
		}

	}
}

func (b *TelegramBot) HandleText(text string, chatID int64) {
	expr := text
	ans, err := arithmetic.SolveExpr(expr)
	slog.Debug("Data", "expr", expr, "ans", ans)

	msg := tapi.NewMessage(chatID, "")
	if err != nil {
		msg.Text = err.Error()
	} else {
		msg.Text = strconv.Itoa(ans)
	}

	_, err = b.BotAPI.Send(msg)
	if err != nil {
		slog.Error("Failed to send message", "error", err, "chat_id", chatID, "message", msg.Text)
	}
}

func (b *TelegramBot) HandleCommand(command string, chatID int64) {
	slog.Debug("Received command", "command", command, "chat_id", chatID)

	var msg tapi.MessageConfig
	switch command {
	case "start":
		msg = tapi.NewMessage(chatID, "Welcome to the Telegram Arithmetic Bot!\n"+
			"Send me a mathematical expression( e.g (43 + 27) * (5 - 3) ) to solve it.\n"+
			"I can support (+), (-), (*), (/), operations and parentheses.")
	default:
		msg = tapi.NewMessage(chatID, fmt.Sprintf("invalid command(/%s)", command))
	}

	_, err := b.BotAPI.Send(msg)
	if err != nil {
		slog.Error("Failed to send start message", "error", err, "chat_id", chatID)
	}

}
