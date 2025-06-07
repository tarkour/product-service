package telegrambot

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/mock"
)

type MockBotApi struct {
	mock.Mock
}

func (m *MockBotApi) Send(c tg.Chattable) (tg.Message, error) {
	args := m.Called(c)
	return args.Get(0).(tg.Message), args.Error(1)
}

func (m *MockBotApi) GetUpdatesChan(config tg.UpdateConfig) tg.UpdatesChannel {
	args := m.Called(config)
	return args.Get(0).(tg.UpdatesChannel)
}
