-- +goose Up
CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_follows;