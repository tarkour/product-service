-- +goose Up
-- +goose StatementBegin
CREATE TABLE type (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255) NOT NULL
);

COMMENT ON TABLE type IS 'Таблица для хранения информации о типе продукта';
COMMENT ON COLUMN type.id IS 'Уникальный идентификатор продукта (первичный ключ)';
COMMENT ON COLUMN type.type IS 'Тип продукта';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS type;
-- +goose StatementEnd
