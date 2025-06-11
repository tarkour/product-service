package telegrambot

import (
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	db "github.com/tarkour/product-service/pkg/database"
)

func NewBotHandler(bot BotAPI, qe db.Executor, adminID int64) *BotHandler {
	return &BotHandler{
		bot:       bot,
		queryExec: qe,
		adminID:   adminID,
	}
}

func (h *BotHandler) HandleQueryCommand(update tg.Update) {
	if update.Message == nil || update.Message.From == nil {
		return
	}

	if update.Message.From.ID != h.adminID {
		msg := tg.NewMessage(update.Message.Chat.ID, "🚫 Access denied")
		h.bot.Send(msg)
		return
	}

	query := strings.TrimPrefix(update.Message.Text, "/query ")
	if query == "" {
		msg := tg.NewMessage(update.Message.Chat.ID, "❌ Please provide a query after /query command")
		h.bot.Send(msg)
		return
	}

	result, err := h.queryExec.Execute(query)

	msg := tg.NewMessage(update.Message.Chat.ID, "")
	if err != nil {
		msg.Text = "❌ Error: " + escapeMarkdown(err.Error())
	} else {
		msg.Text = "📊 Result:\n\n" + escapeMarkdown(result)
		msg.ParseMode = "MarkdownV2"
	}

	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("failed to send result message: %v", err)
	}
}

func escapeMarkdown(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(text)
}

func (h *BotHandler) SendMainMenu(chatID int64) {

	keyboard := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Выгрузка брендов", "download_brands"),
			tg.NewInlineKeyboardButtonData("Товары в наличии", "products_in_stock"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Статистика", "stats"),
			tg.NewInlineKeyboardButtonData("Проданные товары", "sold_products"),
		),
	)

	msg := tg.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = keyboard
	h.bot.Send(msg)

}

func (h *BotHandler) HandleCallbackQuery(update tg.Update) {

	callback := update.CallbackQuery
	if callback == nil {
		return
	}

	callbackConf := tg.NewCallback(callback.ID, "")
	if _, err := h.bot.Request(callbackConf); err != nil {
		log.Printf("Callback confirmation error: %v", err)
	}

	switch callback.Data {
	case "download_brands":
		h.handleBrandsDownload(callback.Message.Chat.ID)
		// case "products_in_stock":
		// 	h.handleProductsInStock(callback.Message.Chat.ID)
		// case "stats":
		// 	h.handleStats(callback.Message.Chat.ID)
		// case "sold_products":
		// 	h.handleSoldProducts(callback.Message.Chat.ID)
	}
}

func (h *BotHandler) handleBrandsDownload(chatID int64) {

	msg := tg.NewMessage(chatID, "⏳ Загружаю данные о брендах...")

	sentMsg, _ := h.bot.Send(msg)

	result, err := h.queryExec.Execute("SELECT id, name FROM brand;")

	deleteMsg := tg.NewDeleteMessage(chatID, sentMsg.MessageID)
	h.bot.Send(deleteMsg)

	if err != nil {
		h.bot.Send(tg.NewMessage(chatID, "❌ Ошибка: "+err.Error()))
		return
	}

	response := tg.NewMessage(chatID, "Список брендов:\n\n"+escapeMarkdown(result))
	response.ParseMode = "MarkdownV2"
	h.bot.Send(response)
}
