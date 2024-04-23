// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: post.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (created_at,updated_at,title, content, number_phone, address, nick_name, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, title, content, number_phone, address, nick_name, user_id
`

type CreatePostParams struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Content     string
	NumberPhone string
	Address     string
	NickName    string
	UserID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Content,
		arg.NumberPhone,
		arg.Address,
		arg.NickName,
		arg.UserID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Content,
		&i.NumberPhone,
		&i.Address,
		&i.NickName,
		&i.UserID,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, created_at, updated_at, title, content, number_phone, address, nick_name, user_id FROM posts
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Content,
			&i.NumberPhone,
			&i.Address,
			&i.NickName,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
