-- +goose Up
-- +goose StatementBegin
BEGIN;

ALTER TABLE walking_record_entities
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE user_entities
DROP COLUMN created_at,
DROP COLUMN updated_at;

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

ALTER TABLE walking_record_entities
DROP COLUMN created_at,
DROP COLUMN updated_at;

ALTER TABLE user_entities
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

COMMIT;

-- +goose StatementEnd
