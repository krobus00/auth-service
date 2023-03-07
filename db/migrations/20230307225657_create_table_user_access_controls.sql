-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_access_controls (
    user_id varchar(36),
    ac_id varchar(36),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_user_ac UNIQUE (user_id, ac_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_ac FOREIGN KEY(ac_id) REFERENCES access_controls(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_access_controls;
-- +goose StatementEnd
