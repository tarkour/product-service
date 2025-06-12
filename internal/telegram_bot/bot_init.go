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

