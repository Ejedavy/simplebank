-- name: CreateEntry :one
INSERT INTO "Entry" (
    account_id, amount
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetEntry :one
SELECT * FROM "Entry"
WHERE id = $1;


-- name: ListBankEntries :many
SELECT * FROM "Entry"
ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListAccountEntries :many
SELECT * FROM "Entry"
WHERE account_id = $1
ORDER BY id LIMIT $2 OFFSET $3;


