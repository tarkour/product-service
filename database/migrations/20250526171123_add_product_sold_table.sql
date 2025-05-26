-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_sold (
    id SERIAL PRIMARY KEY,
    brand_id INT NOT NULL REFERENCES brand(id),
    type_id INT NOT NULL REFERENCES type(id),
    color_id INT NOT NULL REFERENCES color(id),
    created_at TIMESTAMP NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    sold_at TIMESTAMP DEFAULT NOW() NOT NULL
);

COMMENT ON TABLE product_sold IS 'Проданные уникальные товары';
COMMENT ON COLUMN product_sold.brand_id IS 'Внешний ключ на таблицу brand';
COMMENT ON COLUMN product_sold.type_id IS 'Внешний ключ на таблицу type';
COMMENT ON COLUMN product_sold.color_id IS 'Внешний ключ на таблицу color';
COMMENT ON COLUMN product_sold.created_at IS 'Дата добавления в product_in_stock';
COMMENT ON COLUMN product_sold.sold_at IS 'Время продажи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_sold;
-- +goose StatementEnd
