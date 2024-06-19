-- +goose Up
-- +goose StatementBegin
BEGIN;

ALTER TABLE user_entities
ADD CONSTRAINT unique_user_entities_auth_id UNIQUE (auth_id);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

ALTER TABLE user_entities
DROP CONSTRAINT unique_user_entities_auth_id;

COMMIT;

-- +goose StatementEnd
