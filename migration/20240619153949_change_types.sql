-- +goose Up
-- +goose StatementBegin
BEGIN;

DROP TABLE walking_record_entities;

DROP TABLE follow_entities;

ALTER TABLE user_entities
DROP CONSTRAINT user_entities_pkey,
DROP CONSTRAINT unique_user_entities_user_id;

ALTER TABLE user_entities
RENAME COLUMN id TO auth_id;

ALTER TABLE user_entities
ADD CONSTRAINT user_entities_pkey PRIMARY KEY (user_id);

CREATE TABLE walking_record_entities (
    user_id VARCHAR(255),
    date DATE,
    distance INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, date),
    CONSTRAINT fk_walking_record_entities_user_id FOREIGN KEY (user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE
);

CREATE TABLE follow_entities (
    user_id VARCHAR(255),
    followed_user_id VARCHAR(255),
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

DROP TABLE walking_record_entities;

ALTER TABLE user_entities
DROP CONSTRAINT user_entities_pkey;

ALTER TABLE user_entities
RENAME COLUMN auth_id TO id;

ALTER TABLE user_entities
ADD CONSTRAINT user_entities_pkey PRIMARY KEY (id),
ADD CONSTRAINT unique_user_entities_user_id UNIQUE (user_id);

CREATE TABLE walking_record_entities (
    user_id VARCHAR(255),
    date TIMESTAMP WITH TIME ZONE,
    distance INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, date),
    CONSTRAINT fk_walking_record_entities_user_id FOREIGN KEY (user_id) REFERENCES user_entities (id) ON DELETE CASCADE
);

CREATE TABLE follow_entities (
    user_id VARCHAR(255),
    followed_user_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, followed_user_id),
    CONSTRAINT fk_follow_entities_user_id FOREIGN KEY (user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_follow_entities_followed_user_id FOREIGN KEY (followed_user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE
);

COMMIT;

-- +goose StatementEnd
