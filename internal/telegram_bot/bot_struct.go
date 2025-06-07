package telegrambot

import (
	db "github.com/tarkour/product-service/pkg/database"
)

type BotHandler struct {
	bot       BotAPI
	queryExec db.Executor
	adminID   int64
}
