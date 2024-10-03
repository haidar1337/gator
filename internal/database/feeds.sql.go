// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds(id, feed_name, feed_url, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)

RETURNING id, user_id, feed_name, feed_url, created_at, updated_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	FeedName  string
	FeedUrl   string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.FeedName,
		arg.FeedUrl,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedName,
		&i.FeedUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeedsWithUsers = `-- name: GetFeedsWithUsers :many
SELECT feed_name, feed_url, users.name FROM feeds
INNER JOIN users
ON users.id = feeds.user_id
`

type GetFeedsWithUsersRow struct {
	FeedName string
	FeedUrl  string
	Name     string
}

func (q *Queries) GetFeedsWithUsers(ctx context.Context) ([]GetFeedsWithUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsWithUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedsWithUsersRow
	for rows.Next() {
		var i GetFeedsWithUsersRow
		if err := rows.Scan(&i.FeedName, &i.FeedUrl, &i.Name); err != nil {
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
