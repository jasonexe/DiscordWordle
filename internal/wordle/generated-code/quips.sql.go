// Code generated by sqlc. DO NOT EDIT.
// source: quips.sql

package wordle

import (
	"context"
	"database/sql"
)

const createQuipForScore = `-- name: CreateQuipForScore :one
insert into quips (score_value, quip, inside_joke, inside_joke_server_id, created_by_account)
VALUES ($1, $2, $3, $4, $5)
returning id, score_value, quip, inside_joke, inside_joke_server_id, created_by_account, created_at, uses
`

type CreateQuipForScoreParams struct {
	ScoreValue         int32          `json:"score_value"`
	Quip               string         `json:"quip"`
	InsideJoke         bool           `json:"inside_joke"`
	InsideJokeServerID sql.NullString `json:"inside_joke_server_id"`
	CreatedByAccount   string         `json:"created_by_account"`
}

func (q *Queries) CreateQuipForScore(ctx context.Context, arg CreateQuipForScoreParams) (Quip, error) {
	row := q.queryRow(ctx, q.createQuipForScoreStmt, createQuipForScore,
		arg.ScoreValue,
		arg.Quip,
		arg.InsideJoke,
		arg.InsideJokeServerID,
		arg.CreatedByAccount,
	)
	var i Quip
	err := row.Scan(
		&i.ID,
		&i.ScoreValue,
		&i.Quip,
		&i.InsideJoke,
		&i.InsideJokeServerID,
		&i.CreatedByAccount,
		&i.CreatedAt,
		&i.Uses,
	)
	return i, err
}

const deleteQuipByIdAndServerId = `-- name: DeleteQuipByIdAndServerId :exec
delete
from quips
where id = $1
  and inside_joke_server_id = $2
`

type DeleteQuipByIdAndServerIdParams struct {
	ID                 int64          `json:"id"`
	InsideJokeServerID sql.NullString `json:"inside_joke_server_id"`
}

func (q *Queries) DeleteQuipByIdAndServerId(ctx context.Context, arg DeleteQuipByIdAndServerIdParams) error {
	_, err := q.exec(ctx, q.deleteQuipByIdAndServerIdStmt, deleteQuipByIdAndServerId, arg.ID, arg.InsideJokeServerID)
	return err
}

const getQuipByScore = `-- name: GetQuipByScore :one
SELECT id, score_value, quip, inside_joke, inside_joke_server_id, created_by_account, created_at, uses
FROM quips
where score_value = $1
  and (not inside_joke or (inside_joke and inside_joke_server_id = $2))
ORDER BY uses, random()
LIMIT 1
`

type GetQuipByScoreParams struct {
	ScoreValue         int32          `json:"score_value"`
	InsideJokeServerID sql.NullString `json:"inside_joke_server_id"`
}

func (q *Queries) GetQuipByScore(ctx context.Context, arg GetQuipByScoreParams) (Quip, error) {
	row := q.queryRow(ctx, q.getQuipByScoreStmt, getQuipByScore, arg.ScoreValue, arg.InsideJokeServerID)
	var i Quip
	err := row.Scan(
		&i.ID,
		&i.ScoreValue,
		&i.Quip,
		&i.InsideJoke,
		&i.InsideJokeServerID,
		&i.CreatedByAccount,
		&i.CreatedAt,
		&i.Uses,
	)
	return i, err
}

const getQuipsByCreatedByAccount = `-- name: GetQuipsByCreatedByAccount :many
SELECT id, score_value, quip, inside_joke, inside_joke_server_id, created_by_account, created_at, uses
FROM quips
where created_by_account = $1
`

func (q *Queries) GetQuipsByCreatedByAccount(ctx context.Context, createdByAccount string) ([]Quip, error) {
	rows, err := q.query(ctx, q.getQuipsByCreatedByAccountStmt, getQuipsByCreatedByAccount, createdByAccount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quip
	for rows.Next() {
		var i Quip
		if err := rows.Scan(
			&i.ID,
			&i.ScoreValue,
			&i.Quip,
			&i.InsideJoke,
			&i.InsideJokeServerID,
			&i.CreatedByAccount,
			&i.CreatedAt,
			&i.Uses,
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

const getQuipsByServerId = `-- name: GetQuipsByServerId :many
select id, score_value, quip, inside_joke, inside_joke_server_id, created_by_account, created_at, uses
from quips
where inside_joke_server_id = $1
order by score_value, id
`

func (q *Queries) GetQuipsByServerId(ctx context.Context, insideJokeServerID sql.NullString) ([]Quip, error) {
	rows, err := q.query(ctx, q.getQuipsByServerIdStmt, getQuipsByServerId, insideJokeServerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quip
	for rows.Next() {
		var i Quip
		if err := rows.Scan(
			&i.ID,
			&i.ScoreValue,
			&i.Quip,
			&i.InsideJoke,
			&i.InsideJokeServerID,
			&i.CreatedByAccount,
			&i.CreatedAt,
			&i.Uses,
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

const incrementQuip = `-- name: IncrementQuip :exec
UPDATE quips
SET uses = uses + 1
WHERE id = $1
`

func (q *Queries) IncrementQuip(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.incrementQuipStmt, incrementQuip, id)
	return err
}
