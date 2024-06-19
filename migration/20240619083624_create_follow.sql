-- +goose Up
-- +goose StatementBegin
BEGIN;

DROP INDEX user_entities_user_id_idx;

ALTER TABLE user_entities
ADD CONSTRAINT unique_user_entities_user_id UNIQUE (user_id);

ALTER TABLE walking_record_entities
DROP CONSTRAINT fk_walking_record_entities_user_id;

ALTER TABLE walking_record_entities
ADD FOREIGN KEY fk_walking_record_entities_user_id (user_id) REFERENCES user_entities (user_id);

CREATE TABLE follow_entities (
    user_id VARCHAR(36) NOT NULL,
    followed_user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, followed_user_id),
    CONSTRAINT fk_follow_entities_user_id FOREIGN KEY (user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_follow_entities_followed_user_id FOREIGN KEY (followed_user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE
);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE follow_entities;

ALTER TABLE walking_record_entities
DROP CONSTRAINT fk_walking_record_entities_user_id;

ALTER TABLE walking_record_entities
ADD FOREIGN KEY fk_walking_record_entities_user_id (user_id) REFERENCES user_entities (id);

ALTER TABLE user_entities
DROP CONSTRAINT unique_user_entities_user_id;

CREATE INDEX user_entities_user_id_idx ON user_entities (user_id);

COMMIT;

-- +goose StatementEnd
