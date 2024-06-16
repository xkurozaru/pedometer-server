-- +goose Up
-- +goose StatementBegin
BEGIN;

ALTER TABLE user_entities
ALTER COLUMN id TYPE VARCHAR(36);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

ALTER TABLE user_entities
ALTER COLUMN id TYPE VARCHAR(28);

COMMIT;

-- +goose StatementEnd
