-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todo(
id serial primary key,
name text not null,
description text not null,
user_id integer not null,
foreign key (user_id) references users(id) on delete cascade
);
CREATE INDEX idx_todo_name on todo(name);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo;
DROP INDEX IF EXISTS idx_todo_name;
-- +goose StatementEnd
