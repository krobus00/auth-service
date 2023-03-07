-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_controls (
    id varchar(36) UNIQUE,
	name varchar(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_controls;
-- +goose StatementEnd
