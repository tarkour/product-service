-- +goose Up
-- +goose StatementBegin
CREATE TABLE color (
    id SERIAL PRIMARY KEY,
    color VARCHAR(255) NOT NULL
);

COMMENT ON TABLE color IS 'Таблица для хранения информации о цвете продукта';
COMMENT ON COLUMN color.id IS 'Уникальный идентификатор продукта (первичный ключ)';
COMMENT ON COLUMN color.color IS 'Цвет продукта';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS color;
-- +goose StatementEnd
