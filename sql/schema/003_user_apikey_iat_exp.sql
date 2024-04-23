-- +goose Up
ALTER TABLE users 
    ADD COLUMN api_iat TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN api_exp TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '1 hour';

-- +goose Down
ALTER TABLE users DROP COLUMN api_iat, DROP COLUMN api_exp;
