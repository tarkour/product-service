-- +goose Up
-- +goose StatementBegin
CREATE TABLE test_table(
    id SERIAL PRIMARY KEY
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS test_table;
-- +goose StatementEnd
