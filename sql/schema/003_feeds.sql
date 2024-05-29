-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;