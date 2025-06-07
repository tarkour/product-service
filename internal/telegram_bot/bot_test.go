package telegrambot_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/mock"
	tgbot "github.com/tarkour/product-service/internal/telegram_bot"
	db "github.com/tarkour/product-service/pkg/database"
)

func TestHanldeQueryCommand(t *testing.T) {

	mockBot := new(tgbot.MockBotApi)
	mockQueryExec := new(db.MockQueryExecutor)

	handler := tgbot.NewBotHandler(mockBot, mockQueryExec, 123)

	t.Run("Valid query", func(t *testing.T) {

		mockQueryExec.On(
			"Execute",
			"SELECT * FROM brands LIMIT 1",
		).Return(
			"id | name\n1 | Nike",
			nil,
		)

		mockBot.On("Send", mock.AnythingOfType("tg.MessageConfig")).
			Return(tg.Message{}, nil)

		handler.HandleQueryCommand(tg.Update{
			Message: &tg.Message{
				MessageID: 1,
				From:      &tg.User{ID: 123},
				Chat:      &tg.Chat{ID: 1},
				Text:      "/query SELECT * FROM brand LIMIT 1",
			}})

		mockQueryExec.AssertExpectations(t)
		mockBot.AssertExpectations(t)

	})

	t.Run("Dangerous query", func(t *testing.T) {
		mockQueryExec.On(
			"Execute",
			"DROP TABLE products",
		).Return(
			"",
			fmt.Errorf("dangerous qeury blocked"),
		)
		mockBot.On(
			"Send",
			mock.MatchedBy(func(c tg.Chattable) bool {
				msg := c.(tg.MessageConfig)
				return strings.Contains(msg.Text, "dangerous query blocked")
			}),
		).Return(
			tg.Message{},
			nil,
		)

		msg := &tg.Message{
			Text: "/query DROP TABLE brand",
			From: &tg.User{ID: 123},
		}

		handler.HandleQueryCommand(tg.Update{Message: msg})
		mockBot.AssertCalled(t, "Send", mock.Anything)
	})

	t.Run("GetMe info Test Success", func(t *testing.T) {
		mockBot.On("GetMe").Return(tg.User{
			ID:       123456789,
			UserName: "test_bot",
		}, nil) //success
	})

	t.Run("GetMe info Test Error", func(t *testing.T) {
		mockBot.On("GetMe info Test Error").Return(tg.User{}, errors.New("API timeout")) //error
	})

}

// func TestBotDatabaseIntegration(t *testing.T) {

// 	// cfg, err := conf.ReadConfig()
// 	// if err != nil {
// 	// 	log.Print("Config error: ", err)
// 	// }

// 	conn, err := db.ConnectDB(path)
// 	assert.NoError(t, err, "DB should succeed")
// 	defer conn.Close(context.Background())

// 	queryExec := db.NewQueryExecutor(conn, true, slog.Default())

// 	bot, err := tgbotapi.NewBotAPI("TEST_TOKEN")
// 	assert.NoError(t, err, "bot init should succeed")

// 	testCases := []struct {
// 		name     string
// 		query    string
// 		expected string
// 		hasError bool
// 	}{
// 		{
// 			name:     "Valid SELECT",
// 			query:    "SELECT * FROM brand LIMIT 1;",
// 			expected: "id | brand",
// 			hasError: false,
// 		},
// 		{
// 			name:     "Invalid query",
// 			query:    "SELECT * FROM non_exictent_table;",
// 			hasError: true,
// 		},
// 		{
// 			name:     "Dangerous query blocked",
// 			query:    "DROP TABLE test_table;",
// 			hasError: true,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			result, err := queryExec.Execute(tc.query)

// 			if tc.hasError {
// 				assert.Error(t, err, "Expected Error")
// 			} else {
// 				assert.NoError(t, err, "Expected No Error")
// 				assert.Contains(t, result, tc.expected, "Result should contain expected data")
// 			}
// 		})
// 	}

// 	t.Run("Bot command handling", func(t *testing.T) {
// 		msg := tgbotapi.Message{
// 			Text: "/query SELECT * FROM brand",
// 			From: &tgbotapi.User{ID: 123},
// 		}

// 		//handeling msg

// 		_, err := queryExec.Execute("SELECT * FROM brand")
// 		assert.NoError(t, err, "Query should execute successfully")
// 	})

// }
