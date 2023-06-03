-- name: CreateAccountType :one
INSERT INTO account_type (id, name) VALUES (NULL, ?) RETURNING *;

-- name: GetAccountType :one
SELECT id, name FROM account_type WHERE name = ?;

