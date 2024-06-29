-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_lists;
-- +goose StatementEnd
