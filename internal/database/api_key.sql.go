// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: api_key.sql

package database

import (
	"context"
)

const getAPIKeyByEmail = `-- name: GetAPIKeyByEmail :one
SELECT api_key FROM users WHERE email = $1
`

func (q *Queries) GetAPIKeyByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, getAPIKeyByEmail, email)
	var api_key string
	err := row.Scan(&api_key)
	return api_key, err
}
