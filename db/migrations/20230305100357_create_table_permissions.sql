-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id varchar(36) UNIQUE,
	name varchar(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
