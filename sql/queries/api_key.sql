-- name: GetAPIKeyByEmail :one
SELECT api_key FROM users WHERE email = $1;