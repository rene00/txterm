-- name: CreateAccount :one
INSERT INTO account (id, name, description, account_type_id) VALUES (NULL, ?, ?, ?) RETURNING *;

-- name: GetAccount :one
SELECT a.id AS id, a.name AS name, a.description AS description, at.id AS account_type_id, at.name AS account_type_name
FROM account a LEFT JOIN account_type at
WHERE a.name = ? 
AND at.name = ?;

-- name: ListAccounts :many
SELECT a.id AS id, a.name AS name, a.description AS description, at.id AS account_type_id, at.name AS account_type_name
FROM account a LEFT JOIN account_type at
ON a.account_type_id = at.id
ORDER BY a.name;
