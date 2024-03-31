-- name: CreateUser :one
INSERT INTO users(id, first_name, last_name, email, nick_name, number_phone, day_of_birth, address , role, create_at, update_at, api_key)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,
encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;