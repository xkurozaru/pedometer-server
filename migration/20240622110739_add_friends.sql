-- +goose Up
-- +goose StatementBegin
BEGIN;

DROP TABLE follow_entities;

CREATE TABLE friend_entities (
    user_id VARCHAR(255),
    friend_user_id VARCHAR(255),
    status VARCHAR(36),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_user_id),
    CONSTRAINT fk_friend_entities_user_id FOREIGN KEY (user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE,
    CONSTRAINT fk_friend_entities_friend_user_id FOREIGN KEY (friend_user_id) REFERENCES user_entities (user_id) ON DELETE CASCADE
);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE friend_entities;

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
