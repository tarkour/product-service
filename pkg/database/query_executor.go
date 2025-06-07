package db

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// NewQueryExecutor создает экземпляр QueryExecutor
func NewQueryExecutor(conn *pgx.Conn, safeMode bool, logger *slog.Logger) *QueryExecutor {
	return &QueryExecutor{
		conn:     conn,
		safeMode: safeMode,
		logger:   logger,
	}
}

// Execute выполняет SQL-запрос и возвращает результат в виде таблицы (Markdown)
func (qe *QueryExecutor) Execute(query string) (string, error) {
	// Проверка безопасности запроса
	if qe.safeMode && !qe.isQuerySafe(query) {
		return "", fmt.Errorf("запрещенный тип запроса")
	}

	// Логирование операции
	qe.logger.Debug("Executing query", "query", query)

	// Выполнение запроса
	rows, err := qe.conn.Query(context.Background(), query)
	if err != nil {
		qe.logger.Error("Query failed", "error", err)
		return "", fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	// Получение метаданных о колонках
	columns := rows.FieldDescriptions()
	headers := make([]string, len(columns))
	for i, col := range columns {
		headers[i] = string(col.Name)
	}

	// Сбор результатов
	var results []string
	results = append(results, strings.Join(headers, " | "))
	results = append(results, strings.Repeat("---", len(headers)))

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			qe.logger.Error("Failed to read row", "error", err)
			return "", fmt.Errorf("ошибка чтения данных: %w", err)
		}

		row := make([]string, len(values))
		for i, val := range values {
			row[i] = qe.formatValue(val)
		}
		results = append(results, strings.Join(row, " | "))
	}

	return strings.Join(results, "\n"), nil
}

// isQuerySafe проверяет запрос на наличие опасных операций
func (qe *QueryExecutor) isQuerySafe(query string) bool {
	forbiddenKeywords := []string{
		"DROP", "DELETE", "TRUNCATE",
		"GRANT", "REVOKE", "ALTER",
	}

	upperQuery := strings.ToUpper(query)
	for _, keyword := range forbiddenKeywords {
		if strings.Contains(upperQuery, keyword) {
			qe.logger.Warn("Blocked dangerous query", "keyword", keyword)
			return false
		}
	}
	return true
}

// formatValue преобразует значения БД в строки
func (qe *QueryExecutor) formatValue(val interface{}) string {
	switch v := val.(type) {
	case pgtype.Numeric:
		num, _ := v.Int64Value()
		return fmt.Sprintf("%d", num.Int64)
	case pgtype.Timestamp:
		return v.Time.Format("2006-01-02 15:04:05")
	default:
		return fmt.Sprintf("%v", v)
	}
}
