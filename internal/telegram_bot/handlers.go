package telegrambot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tarkour/product-service/pkg/database"
)

type BotHandler struct {
	bot       *tgbotapi.BotAPI
	queryExec *database.QueryExecutor
	adminID   int64
}

func NewBotHandler(bot *tgbotapi.BotAPI, qe *database.QueryExecutor, adminID int64) *BotHandler {
	return &BotHandler{
		bot:       bot,
		queryExec: qe,
		adminID:   adminID,
	}

}

func (h *BotHandler) HandleQueryCommand(update tgbotapi.Update) {
	if update.Message.From.ID != h.adminID {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🚫 Access denied")
		h.bot.Send(msg)
		return
	}

	query := strings.TrimPrefix(update.Message.Text, "/query ")
	result, err := h.queryExec.Execute(query)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if err != nil {
		msg.Text = "❌ Error: " + err.Error()
	} else {
		msg.Text = "📊 Result:\n```\n" + result + "\n```"
		msg.ParseMode = "MarkdownV2"
	}

	h.bot.Send(msg)
}
