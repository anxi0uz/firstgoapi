-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
id serial primary key,
name text not null
);
CREATE INDEX idx_users_name on users(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS idx_users_name;
-- +goose StatementEnd
