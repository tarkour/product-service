package db

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

// QueryExecutor инкапсулирует логику выполнения SQL-запросов
type QueryExecutor struct {
	conn     *pgx.Conn    // Подключение к БД
	safeMode bool         // Флаг для блокировки опасных запросов
	logger   *slog.Logger // Логгер
}

type Executor interface {
	Execute(query string) (string, error)
}
