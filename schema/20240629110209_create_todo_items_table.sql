-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_items
-- +goose StatementEnd
