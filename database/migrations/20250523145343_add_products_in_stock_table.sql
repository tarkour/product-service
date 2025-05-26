-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_in_stock (
    id SERIAL PRIMARY KEY,
    brand_id INT NOT NULL REFERENCES brand(id),
    type_id INT NOT NULL REFERENCES type(id),
    color_id INT NOT NULL REFERENCES color(id),
    created_at TIMESTAMP DEFAULT NOW(),
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT DEFAULT 0
);

COMMENT ON TABLE product_in_stock IS 'Товары в наличии';
COMMENT ON COLUMN product_in_stock.brand_id IS 'Внешний ключ на таблицу brand';
COMMENT ON COLUMN product_in_stock.type_id IS 'Внешний ключ на таблицу type';
COMMENT ON COLUMN product_in_stock.color_id IS 'Внешний ключ на таблицу color';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_in_stock;
-- +goose StatementEnd
