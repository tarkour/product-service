-- +goose Up
-- +goose StatementBegin
CREATE TABLE brand (
    id SERIAL PRIMARY KEY,
    brand VARCHAR(255) NOT NULL
);

COMMENT ON TABLE brand IS 'Таблица для хранения информации о брендах продуктов';
COMMENT ON COLUMN brand.id IS 'Уникальный идентификатор продукта (первичный ключ)';
COMMENT ON COLUMN brand.brand IS 'Название бренда продукта';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brand;
-- +goose StatementEnd
