-- name: CreateLedgerEntry :one
INSERT INTO ledger_entries (user_id, amount, type)
VALUES ($1, $2, $3)
RETURNING *;