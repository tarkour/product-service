package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbot "github.com/tarkour/product-service/internal/telegram_bot"
	"github.com/tarkour/product-service/pkg/config"
	db "github.com/tarkour/product-service/pkg/database"
)

const (
	path = "./internal/config"
)

func main() {

	cfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	conn, err := db.ConnectDB(path)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	fmt.Println("Database connected successfully")

	queryExec := db.NewQueryExecutor(conn, cfg.Telegram.Safe_mode, slog.Default())

	//tgbot launch
	bot, err := tgbot.InitBot(cfg.Telegram.Token)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	botHandler := tgbot.NewBotHandler(bot, queryExec, cfg.Telegram.Admin_ID)

	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "query":
				botHandler.HandleQueryCommand(update)
			}
		}

	}
}
