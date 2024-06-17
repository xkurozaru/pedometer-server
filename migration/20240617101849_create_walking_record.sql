-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE TABLE walking_record_entities (
    id VARCHAR(28) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id VARCHAR(255) NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    steps INTEGER NOT NULL,
    distance INTEGER NOT NULL,
    calories INTEGER NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES user_entities (id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_walking_record_entities_user_id_date ON walking_record_entities (user_id, date);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE walking_record_entities;

COMMIT;

-- +goose StatementEnd
