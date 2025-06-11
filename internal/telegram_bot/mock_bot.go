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

// гарантированно свежие данные о боте
func (m *MockBotApi) GetMe() (tg.User, error) {
	args := m.Called()
	return args.Get(0).(tg.User), args.Error(1)
}

// работает в контексте, при правильной инициализации бота
func (m *MockBotApi) Self() tg.User {
	args := m.Called()
	return args.Get(0).(tg.User)
}

func (m *MockBotApi) Debug() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockBotApi) SetDebug(debug bool) {
	m.Called(debug)
}

func (m *MockBotApi) Request(c tg.Chattable) (*tg.APIResponse, error) {
	args := m.Called(c)
	return args.Get(0).(*tg.APIResponse), args.Error(1)
}
