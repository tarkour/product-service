package telegrambot

import (
	"context"
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
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

	// keyboard := tg.NewInlineKeyboardMarkup(
	// 	tg.NewInlineKeyboardRow(
	// 		h.SendBrandButton(),
	// getMainMenu(),
	// ),
	// tg.NewInlineKeyboardRow(
	// 	tg.NewInlineKeyboardButtonData("Статистика", "stats"),
	// 	tg.NewInlineKeyboardButtonData("Проданные товары", "sold_products"),
	// ),
	// )

	keyboard := h.GetMainMenu()

	msg := tg.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = keyboard
	h.bot.Send(msg)

}

func (h *BotHandler) SendBrandButton() tg.InlineKeyboardButton {
	return tg.NewInlineKeyboardButtonData("Показать бренды", "query:SELECT brand FROM brand;")
}

func (h *BotHandler) CreateProductInStock() tg.InlineKeyboardButton {
	return tg.NewInlineKeyboardButtonData("Добавить товар", "query:")
}

func GetMainMenu() tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(tg.NewKeyboardButton("Создать запись")),
		tg.NewKeyboardButtonRow(tg.NewKeyboardButton("Выберите бренд")),
		tg.NewKeyboardButtonRow(tg.NewKeyboardButton("Выберите тип")),
		tg.NewKeyboardButtonRow(tg.NewKeyboardButton("Выберите цвет")),
		tg.NewKeyboardButtonRow(tg.NewKeyboardButton("Определите цену")),
		tg.NewKeyboardButtonRow(tg.NewKeyboardButton("Определите количество")),
	)
}

var database *pgx.Conn

func (h *BotHandler) GetOptionsFromTable(c context.Context, tablename string) ([]string, error) {
	query := "SELECT name FROM" + tablename + ";"
	rows, err := database.Query(c, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var options []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		options = append(options, name)
	}
	return options, nil
}

func (h *BotHandler) CreateKeyBordFromOptions(options []string) tg.ReplyKeyboardMarkup {
	var buttons []tg.KeyboardButton

	for _, opt := range options {
		buttons = append(buttons, tg.NewKeyboardButton(opt))
	}

	return tg.NewReplyKeyboard(buttons)
}

func (h *BotHandler) HandleMessage(c context.Context, update tg.Update, state *string, tempData *map[string]string) {
	msg := update.Message

	switch msg.Text {
	case "Создать запись":
		reply := tg.NewMessage(msg.Chat.ID, "Выберите действие:")
		reply.ReplyMarkup = h.GetMainMenu()
		h.bot.Send(reply)
		*state = ""
	case "Выберите бренд", "Выберите тип", "Выберите цвет":
		var tableName string
		switch msg.Text {
		case "Выберите бренд":
			tableName = "brand"
		case "Выберите тип":
			tableName = "type"
		case "Выберите цвет":
			tableName = "color"
		}

		options, err := h.GetOptionsFromTable(c, tableName)
		if err != nil {
			log.Printf("Recieving data error: %v", err)
			return
		}

		keyboard := h.CreateKeyBordFromOptions(options)
		reply := tg.NewMessage(msg.Chat.ID, "Выберите "+msg.Text[8:]+":")
		reply.ReplyMarkup = keyboard
		h.bot.Send(reply)
		*state = "wating_for_selection_" + tableName

	case "Определите цену":
		reply := tg.NewMessage(msg.Chat.ID, "Пожалуйста, введите цену (макс. 2 знака после запятой):")
		h.bot.Send(reply)
		*state = "waiting_for_price"
	case "Определите количество":
		reply := tg.NewMessage(msg.Chat.ID, "Пожалуйста, введите количество  (целое число):")
		h.bot.Send(reply)
		*state = "waiting_for_quantity"
	default:
		if *state == "waiting_for_price" {
			//TODO: add check if it's number
			(*tempData)["price"] = msg.Text
			reply := tg.NewMessage(msg.Chat.ID, "Цена сохранена: "+msg.Text)
			h.bot.Send(reply)
			*state = ""
		} else if *state == "waiting_for_quantity" {
			//TODO: add check if it's number without ,.
			(*tempData)["quantity"] = msg.Text
			reply := tg.NewMessage(msg.Chat.ID, "Количество сохранено: "+msg.Text)
			h.bot.Send(reply)
			*state = ""
		} else {
			//TODO: unknown message or command - check? ask about it
			reply := tg.NewMessage(msg.Chat.ID, "Пожалуйста, выберите действие.")
			h.bot.Send(reply)
			reply.ReplyMarkup = h.GetMainMenu()
			h.bot.Send(reply)
		}

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

	h.executeAndSendQuery(update.Message.Chat.ID, query)
}

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
		// TODO: add different butttons (chain button)
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
