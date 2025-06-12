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

func (h *BotHandler) SendMainMenu(chatID int64) {

	keyboard := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Показать бренды", "query:SELECT * FROM brand;"),
			// tg.NewInlineKeyboardButtonData("Товары в наличии", "products_in_stock"),
		),
		// tg.NewInlineKeyboardRow(
		// 	tg.NewInlineKeyboardButtonData("Статистика", "stats"),
		// 	tg.NewInlineKeyboardButtonData("Проданные товары", "sold_products"),
		// ),
	)

	msg := tg.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = keyboard
	h.bot.Send(msg)

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
	// if query == "" {
	// 	msg := tg.NewMessage(update.Message.Chat.ID, "❌ Please provide a query after /query command")
	// 	h.bot.Send(msg)
	// 	return
	// }

	h.executeAndSendQuery(update.Message.Chat.ID, query)
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

// func (h *BotHandler) HandleCallbackQuery(update tg.Update) {

// 	callback := update.CallbackQuery
// 	if callback == nil {
// 		return
// 	}

// 	callbackConf := tg.NewCallback(callback.ID, "")
// 	if _, err := h.bot.Request(callbackConf); err != nil {
// 		log.Printf("Callback confirmation error: %v", err)
// 	}

// 	switch callback.Data {
// 	case "query:SELECT * FROM brand;":
// 		h.HandleButtonPress(update)
// 		// case "products_in_stock":
// 		// 	h.handleProductsInStock(callback.Message.Chat.ID)
// 		// case "stats":
// 		// 	h.handleStats(callback.Message.Chat.ID)
// 		// case "sold_products":
// 		// 	h.handleSoldProducts(callback.Message.Chat.ID)
// 	}
// }

// func (h *BotHandler) handleBrandsDownload(chatID int64) {

// 	msg := tg.NewMessage(chatID, "⏳ Загружаю данные о брендах...")

// 	sentMsg, _ := h.bot.Send(msg)

// 	result, err := h.queryExec.Execute("SELECT * FROM brand;")

// 	deleteMsg := tg.NewDeleteMessage(chatID, sentMsg.MessageID)
// 	h.bot.Send(deleteMsg)

// 	if err != nil {
// 		h.bot.Send(tg.NewMessage(chatID, "❌ Ошибка: "+err.Error()))
// 		return
// 	}

// 	response := tg.NewMessage(chatID, "Список брендов:\n\n"+escapeMarkdown(result))
// 	response.ParseMode = "MarkdownV2"
// 	h.bot.Send(response)
// }

func (h *BotHandler) HandleButtonPress(update tg.Update) {
	callback := update.CallbackQuery
	if callback == nil || callback.Data == "" {
		return
	}

	if callback.From.ID != h.adminID {
		msg := tg.NewMessage(callback.From.ID, "🚫 Access denied")
		h.bot.Send(msg)
		return
	}

	switch {
	case strings.HasPrefix(callback.Data, "query:"):
		query := strings.TrimPrefix(callback.Data, "query:")
		h.executeAndSendQuery(callback.Message.Chat.ID, query)
		// TODO: add different butttons
	}

	callbackCfg := tg.NewCallback(callback.ID, "")
	if _, err := h.bot.Request(callbackCfg); err != nil {
		log.Printf("Callback error: %v", err)
	}
}

func (h *BotHandler) executeAndSendQuery(chatID int64, query string) {
	result, err := h.queryExec.Execute(query)

	msg := tg.NewMessage(chatID, "")
	if err != nil {
		msg.Text = "Error: " + escapeMarkdown(err.Error())
	} else {
		msg.Text = "Result: \n\n" + escapeMarkdown(result)
		msg.ParseMode = "MarkdownV2"
	}

	if _, err := h.bot.Send(msg); err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
