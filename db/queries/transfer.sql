-- name: CreateTransfer :one
INSERT INTO "Transfer" (
    from_account_id, to_account_id, amount
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: GetTransfer :one
SELECT * FROM "Transfer"
WHERE id = $1;

-- name: ListIncomingTransfers :many
SELECT * FROM "Transfer"
WHERE to_account_id = $1
ORDER BY id LIMIT $2 OFFSET $3;

-- name: ListOutgoingTransfers :many
SELECT * FROM "Transfer"
WHERE from_account_id = $1
ORDER BY id LIMIT $2 OFFSET $3;

-- name: ListTransfersBetween :many
SELECT * FROM "Transfer"
WHERE (from_account_id = sqlc.arg(account1) AND to_account_id = sqlc.arg(account2)) OR (from_account_id = sqlc.arg(account2) AND to_account_id = sqlc.arg(account1))
ORDER BY id LIMIT $1 OFFSET $2;


