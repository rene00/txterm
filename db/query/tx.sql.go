// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: tx.sql

package query

import (
	"context"
	"time"
)

const createTx = `-- name: CreateTx :one
INSERT INTO tx (id, date_created, date_posted, memo, amount_num, amount_den, import_id) VALUES (NULL, ?, ?, ?, ?, ?, ?) RETURNING id, date_created, date_posted, memo, amount_num, amount_den, import_id
`

type CreateTxParams struct {
	DateCreated time.Time
	DatePosted  time.Time
	Memo        string
	AmountNum   int64
	AmountDen   int64
	ImportID    int64
}

func (q *Queries) CreateTx(ctx context.Context, arg CreateTxParams) (Tx, error) {
	row := q.db.QueryRowContext(ctx, createTx,
		arg.DateCreated,
		arg.DatePosted,
		arg.Memo,
		arg.AmountNum,
		arg.AmountDen,
		arg.ImportID,
	)
	var i Tx
	err := row.Scan(
		&i.ID,
		&i.DateCreated,
		&i.DatePosted,
		&i.Memo,
		&i.AmountNum,
		&i.AmountDen,
		&i.ImportID,
	)
	return i, err
}

const getDuplicateTx = `-- name: GetDuplicateTx :many
SELECT id, date_created, date_posted, memo, amount_num, amount_den, import_id FROM tx WHERE date_posted = ? AND memo = ? AND amount_num = ? AND amount_den = ?
`

type GetDuplicateTxParams struct {
	DatePosted time.Time
	Memo       string
	AmountNum  int64
	AmountDen  int64
}

func (q *Queries) GetDuplicateTx(ctx context.Context, arg GetDuplicateTxParams) ([]Tx, error) {
	rows, err := q.db.QueryContext(ctx, getDuplicateTx,
		arg.DatePosted,
		arg.Memo,
		arg.AmountNum,
		arg.AmountDen,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tx
	for rows.Next() {
		var i Tx
		if err := rows.Scan(
			&i.ID,
			&i.DateCreated,
			&i.DatePosted,
			&i.Memo,
			&i.AmountNum,
			&i.AmountDen,
			&i.ImportID,
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

const getTxs = `-- name: GetTxs :many
SELECT id, date_created, date_posted, memo, amount_num, amount_den, import_id FROM tx ORDER BY id
`

func (q *Queries) GetTxs(ctx context.Context) ([]Tx, error) {
	rows, err := q.db.QueryContext(ctx, getTxs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tx
	for rows.Next() {
		var i Tx
		if err := rows.Scan(
			&i.ID,
			&i.DateCreated,
			&i.DatePosted,
			&i.Memo,
			&i.AmountNum,
			&i.AmountDen,
			&i.ImportID,
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
