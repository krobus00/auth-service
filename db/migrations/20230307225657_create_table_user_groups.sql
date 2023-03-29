-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_groups (
    user_id varchar(36),
    group_id varchar(36),
    CONSTRAINT unique_user_group UNIQUE (user_id, group_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_group FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_groups;
-- +goose StatementEnd
