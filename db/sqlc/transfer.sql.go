// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO "Transfer" (
    from_account_id, to_account_id, amount
) VALUES (
             $1, $2, $3
         )
    RETURNING id, from_account_id, to_account_id, amount, "createdAt"
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, "createdAt" FROM "Transfer"
WHERE id = $1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listIncomingTransfers = `-- name: ListIncomingTransfers :many
SELECT id, from_account_id, to_account_id, amount, "createdAt" FROM "Transfer"
WHERE to_account_id = $1
ORDER BY id LIMIT $2 OFFSET $3
`

type ListIncomingTransfersParams struct {
	ToAccountID int64 `json:"to_account_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListIncomingTransfers(ctx context.Context, arg ListIncomingTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listIncomingTransfers, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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

const listOutgoingTransfers = `-- name: ListOutgoingTransfers :many
SELECT id, from_account_id, to_account_id, amount, "createdAt" FROM "Transfer"
WHERE from_account_id = $1
ORDER BY id LIMIT $2 OFFSET $3
`

type ListOutgoingTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListOutgoingTransfers(ctx context.Context, arg ListOutgoingTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listOutgoingTransfers, arg.FromAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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

const listTransfersBetween = `-- name: ListTransfersBetween :many
SELECT id, from_account_id, to_account_id, amount, "createdAt" FROM "Transfer"
WHERE (from_account_id = $3 AND to_account_id = $4) OR (from_account_id = $4 AND to_account_id = $3)
ORDER BY id LIMIT $1 OFFSET $2
`

type ListTransfersBetweenParams struct {
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
	Account1 int64 `json:"account1"`
	Account2 int64 `json:"account2"`
}

func (q *Queries) ListTransfersBetween(ctx context.Context, arg ListTransfersBetweenParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersBetween,
		arg.Limit,
		arg.Offset,
		arg.Account1,
		arg.Account2,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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
