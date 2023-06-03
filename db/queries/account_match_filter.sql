-- name: CreateAccountMatchFilter :one
INSERT INTO account_match_filter (id, filter, account_match_id) VALUES (NULL, ?, ?) RETURNING *;
