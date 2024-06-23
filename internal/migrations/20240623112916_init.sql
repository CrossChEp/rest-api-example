-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA user_schema;

CREATE TABLE IF NOT EXISTS user_schema.users(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_schema.users;
DROP SCHEMA IF EXISTS user_schema;
-- +goose StatementEnd
