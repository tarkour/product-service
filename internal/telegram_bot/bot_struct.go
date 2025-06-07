package telegrambot

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	db "github.com/tarkour/product-service/pkg/database"
)

type BotHandler struct {
	bot       BotAPI
	queryExec db.Executor
	adminID   int64
}
