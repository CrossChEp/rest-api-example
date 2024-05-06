-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    email TEXT,
    password TEXT,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
