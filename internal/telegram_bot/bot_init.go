package telegrambot

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Инициализация бота
func initBot(token string) error {
	var err error
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("Error with token: %w", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return nil
}
