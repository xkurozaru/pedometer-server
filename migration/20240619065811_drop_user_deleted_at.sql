-- +goose Up
-- +goose StatementBegin
BEGIN;

ALTER TABLE user_entities
DROP COLUMN deleted_at;

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

ALTER TABLE user_entities
ADD COLUMN deleted_at TIMESTAMP;

COMMIT;

-- +goose StatementEnd
