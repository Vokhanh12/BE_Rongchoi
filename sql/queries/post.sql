-- name: CreatePost :one
INSERT INTO posts (created_at,updated_at,title, content, number_phone, address, nick_name, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts;