// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: story.sql

package sqlc

import (
	"context"
)

const createStory = `-- name: CreateStory :one
INSERT INTO stories (
  user_id, carac_ucs_id, body, public
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, user_id, carac_ucs_id, created_on, update_on, body, public
`

type CreateStoryParams struct {
	UserID     int32  `json:"user_id"`
	CaracUcsID int32  `json:"carac_ucs_id"`
	Body       string `json:"body"`
	Public     int32  `json:"public"`
}

func (q *Queries) CreateStory(ctx context.Context, arg CreateStoryParams) (Story, error) {
	row := q.db.QueryRow(ctx, createStory,
		arg.UserID,
		arg.CaracUcsID,
		arg.Body,
		arg.Public,
	)
	var i Story
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CaracUcsID,
		&i.CreatedOn,
		&i.UpdateOn,
		&i.Body,
		&i.Public,
	)
	return i, err
}

const deleteStory = `-- name: DeleteStory :exec
DELETE FROM stories
WHERE user_id = $1 and carac_ucs_id = $2
`

type DeleteStoryParams struct {
	UserID     int32 `json:"user_id"`
	CaracUcsID int32 `json:"carac_ucs_id"`
}

func (q *Queries) DeleteStory(ctx context.Context, arg DeleteStoryParams) error {
	_, err := q.db.Exec(ctx, deleteStory, arg.UserID, arg.CaracUcsID)
	return err
}

const getStorybyCharacter = `-- name: GetStorybyCharacter :many
SELECT id, user_id, carac_ucs_id, created_on, update_on, body, public FROM stories
WHERE carac_ucs_id = $1
`

func (q *Queries) GetStorybyCharacter(ctx context.Context, caracUcsID int32) ([]Story, error) {
	rows, err := q.db.Query(ctx, getStorybyCharacter, caracUcsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Story
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CaracUcsID,
			&i.CreatedOn,
			&i.UpdateOn,
			&i.Body,
			&i.Public,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStorybyUser = `-- name: GetStorybyUser :many
SELECT id, user_id, carac_ucs_id, created_on, update_on, body, public FROM stories
WHERE user_id = $1
`

func (q *Queries) GetStorybyUser(ctx context.Context, userID int32) ([]Story, error) {
	rows, err := q.db.Query(ctx, getStorybyUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Story
	for rows.Next() {
		var i Story
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CaracUcsID,
			&i.CreatedOn,
			&i.UpdateOn,
			&i.Body,
			&i.Public,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}