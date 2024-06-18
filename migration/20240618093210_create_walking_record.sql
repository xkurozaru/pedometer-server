-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE TABLE walking_record_entities (
    user_id VARCHAR(255) NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    distance INTEGER,
    PRIMARY KEY (user_id, date),
    CONSTRAINT fk_walking_record_entities_user_id FOREIGN KEY (user_id) REFERENCES user_entities (id) ON DELETE CASCADE
);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE walking_record_entities;

COMMIT;

-- +goose StatementEnd
