-- +goose Up
-- +goose StatementBegin
CREATE TABLE item_photo (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES product_in_stock(id),
    file_path VARCHAR(512) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

COMMENT ON TABLE item_photo IS 'Фотографии товаров';
COMMENT ON COLUMN item_photo.product_id IS 'Ссылка на товар';
COMMENT ON COLUMN item_photo.is_primary IS 'Основное фото товара';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_photo;
-- +goose StatementEnd
