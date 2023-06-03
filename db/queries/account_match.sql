-- name: CreateAccountMatch :one
INSERT INTO account_match (id, name, description, account_id) VALUES (NULL, ?, ?, ?) RETURNING *;

-- name: GetAccountMatch :one
SELECT id, name, description, account_id FROM account_match WHERE name = ?;
