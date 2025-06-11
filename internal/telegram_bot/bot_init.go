package telegrambot

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Инициализация бота
func InitBot(token string) (*tg.BotAPI, error) {
	var err error
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("error with token: %w", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account @%s", bot.Self.UserName)
	return bot, nil
}

func StartBotProcessing(bot *tg.BotAPI, handler *BotHandler) {
	u := tg.NewUpdate(0)
	
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates{
		switch{
		case update.Message != nil:
			handler.HandleQueryCommand(update)
		case update.CallbackQuery != nil:
			handler.HandleCallbackQuery(update)
		}
	}
}
