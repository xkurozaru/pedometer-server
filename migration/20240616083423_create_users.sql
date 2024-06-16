-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE TABLE user_entities (
    id VARCHAR(28) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL
);

CREATE INDEX user_entities_user_id_idx ON user_entities (user_id);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE user_entities;

COMMIT;

-- +goose StatementEnd
