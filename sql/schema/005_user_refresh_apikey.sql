-- +goose Up
ALTER TABLE users ADD COLUMN refresh_api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT(
    encode(sha256(random()::text::bytea), 'hex')
);
-- +goose Down
ALTER TABLE users DROP COLUMN refresh_api_key;