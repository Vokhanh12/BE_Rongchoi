-- +goose Up
ALTER TABLE users 
    ADD COLUMN ref_api_iat TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN ref_api_exp TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '1 hour';

-- +goose Down
ALTER TABLE users DROP COLUMN ref_api_iat, DROP COLUMN ref_api_exp;
