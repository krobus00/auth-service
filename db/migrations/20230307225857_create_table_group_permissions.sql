-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS group_permissions (
    group_id varchar(36),
    permission_id varchar(36),
    CONSTRAINT unique_group_permissions UNIQUE (group_id, permission_id),
    CONSTRAINT fk_group FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE,
    CONSTRAINT fk_permission FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS group_permissions;
-- +goose StatementEnd
