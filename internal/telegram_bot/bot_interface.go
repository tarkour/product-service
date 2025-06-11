package telegrambot

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotAPI interface {
	Send(c tg.Chattable) (tg.Message, error)
	GetUpdatesChan(config tg.UpdateConfig) tg.UpdatesChannel
	GetMe() (tg.User, error)
	Request(c tg.Chattable) (*tg.APIResponse, error)
}

var _ BotAPI = (*tg.BotAPI)(nil)
